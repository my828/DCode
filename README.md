# DCode
## Motivation
Programming often requires brainstorming, drawing out algorithms, and visualizing the code. Especially when one coder is trying to explain his or her logic to others, it’s usually helpful to draw it out. The current collaborative coding solutions, however, only offer services such as live editing code and video chatting. The idea is to help visual-learning coders to have a better experience in collaborative coding environment.
## What
The client-facing DCode platform will be similar to Codeshare or Coderpad but it also has a built-in canvas for users to write/draw scratch work. The right half of the screen will be a code editor while the left half would have the empty canvas. Both the code editor and the canvas will support real-time collaboration among multiple users.
## Architecture
## User Stories
<table>
  <tr>
    <th>Priority</th>
    <th>User</th>
    <th>Description</th>
    <th>Implementation</th>
  </tr>
  <tr>
    <td>P0</td>
    <td>As a user</td>
    <td>I want to open a DCode page URL shared by peers.</td>
    <td>Upon receiving a POST request at v1/pages, the server will create and return a new DCode page, generate a unique page URL, insert the new page in the DB, create a new Redis session and opens a  new websocket.</td>
  </tr>
  <tr>
    <td>P0</td>
    <td>As a user</td>
    <td>I want to see real-time updates on my DCode page.</td>
    <td>Upon receiving a GET request to v1/pages/{pageID}, the server opens a websocket connection, adds the user’s unique identifier to the Redis session state and resets the expiration time.</td>
  </tr>
  <tr>
    <td>P0</td>
    <td>As a user</td>
    <td>I want to draw an image on an existing DCode page.</td>
    <td>Upon receiving a GET request to the v1/pages/{pageID} endpoint, the client opens a websocket connection and  continues to render content updates in real-time.</td>
  </tr>
  <tr>
    <td>P0</td>
    <td>As a user</td>
    <td>I want to clear the canvas on an existing DCode page.</td>
    <td>Upon receiving a PATCH request to the v1/pages/{pageID}, the server updates the page content in the database, and it pushes a new message to the RabbitMQ.</td>
  </tr>
  <tr>
    <td>P1</td>
    <td>As a user</td>
    <td>I would like to choose a programming language for my DCode page.</td>
    <td>Upon receiving a DELETE request, the server will remove all figures in the database and RabbitMQ.</td>
  </tr>
  <tr>
    <td>P1</td>
    <td>As a user</td>
    <td>I would like to have a syntax highlighter for my DCode page.</td>
    <td>Upon receiving a PATCH request to v1/pages/{pageID}/settings endpoint, the server stores the settings in the database.</td>
  </tr>
</table>


## API Reference
**POST /v1/pages**

Creates a new page with a unique URL and responds with the DCode page object. The request body is type `application/json`
- **201**: Successfully created page.
- **404**: DCode page with unique ID does not exist.

**GET v1/pages/{pageID}**

Request includes the pageID as a query parameter `id`; Responds with figures and code associated with the specific page ID encoded as `application/json`.
- **200**: status OK.
- **404**: page not found.

**PATCH v1/pages/{pageID}/canvas**

Updates the DCode page with the given ID. The request body should be of type `application/json`, with an `edit` field containing the new `figures` coordinates.
- **200**: Successfully updated canvas.
- **404**: Page no longer exists.

**DELETE v1/pages/{pageID}/canvas**

Deletes the DCode canvas contents at pageID from the database and responds with a status of 200 and a plain text message: `canvas cleared`
- **200**: Successfully deleted canvas contents.
- **404**: Page no longer exists

**PATCH v1/pages/{pageID}/editor**

Updates existing editor in mongo page object
- **200**: Successfully applied changes to editor.
- **404**: Page no longer exists
Responds with updated text

## MongoDB Models
### Page Schema

>{ \_id: Schema.Types.ObjectID,
Figures: [],
Code: [],  
createdAt: Time,  
lastEdited: Time,  
}
>

## Redis Store
### Session Store
>type SessionStore struct {  
	Client,  
	SessionDuration,
}
### Session State
>type SessionState struct {  
  SessionBegin,
  users,
}

### Redis Key Value Pairs
**Key:** Randomly generated hash which forms the unique link for the page.  
**Value:**
