"use strict";

/// Entry point for the Pages microservice for DCode

const express = require("express");
const mongoose = require("mongoose");
const handlers = require("./handlers");

// Read in all environment variables
const MONGODB_ADDRESS = `mongodb://${process.env.MONGODB_ADDRESS}/pagesInfo` || `mongodb://localhost:27017/pagesInfo`;
const PAGES_ADDRESS = process.env.PAGES_ADDRESS || `localhost:5000`;

// Connect to MongoDB
mongoose.connect(MONGODB_ADDRESS, {useNewUrlParser: true})
.then((connection) => {
    console.log("connected to mongodb!");
})
.catch((error) => {
    console.log(`error connecting to mongodb: ${error}`);
});

// Pre-defined resource paths
const resourcePaths = {
    page: "/dcode/v1/{pageID}",
    pageCanvas: "/dcode/v1/{pageID}/canvas",
    pageEditor: "/dcode/v1/{pageID}/editor",
};

const app = express();
const [appHost, appPort] = PAGES_ADDRESS.split(":");

// Add resource handlers
app.use(express.json());
app.use(resourcePaths.page, handlers.pageHandler);
app.use(resourcePaths.pageCanvas, handlers.pageCanvasHandler);
app.use(resourcePaths.pageEditor, handlers.pageEditorHandler);

app.listen(appPort, appHost, () => {
    console.log(`pages server is listening at http://${PAGES_ADDRESS}`);
});