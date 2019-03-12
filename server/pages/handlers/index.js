"use strict";

// Collection of all handlers for the Pages service
const pageHandler = require("./pageHandler");
const pageCanvasHandler = require("./pageCanvasHandler");
const pageEditorHandler = require("./pageEditorHandler");

// Export all handlers
module.exports = {
    pageHandler,
    pageCanvasHandler,
    pageEditorHandler
}