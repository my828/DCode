const mongoose = require('mongoose');
const Schema = mongoose.Schema;

const PageSchema = new Schema({
    pageID: { type: String, required: true },
    figures: { type: Object, required: true },
    code: {type: String, required: true},
    editedAt: { type: Date, required: true }
});

module.exports = mongoose.model('Page', PageSchema);