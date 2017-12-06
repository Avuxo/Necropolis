const express = require('express');
const app = express();

const busboy = require('busboy');
const path = require('path');
const fs = require('fs');

app.use('/packages', express.static(__dirname + '/packages'));
app.use('/public',   express.static(__dirname + '/public'));

app.get('/', (req, res) => {
    res.sendFile(__dirname + "/views/index.html");
});

app.listen('8080', () => {
    console.log("Listening on port 8080");
});

app.post('/upload', (req, res) => {
    let bb = new busboy({headers: req.headers});

    /*handle file uploads*/
    bb.on('file', (fieldname, file, filename, encoding, mime) => {
        let saveFile = path.join('./packages', filename + ".tomb");
        console.log(req.connection.remoteAddress + " | uploading | " + filename);
        file.pipe(fs.createWriteStream(saveFile));
    });

    /*file upload finished*/
    bb.on('finish', () => {
        console.log(req.connection.remoteAddress + " | finished uploading");
    });

    return req.pipe(bb);
});
