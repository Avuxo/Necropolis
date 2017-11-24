const express = require('express');
const app = express();

app.use('/packages', express.static(__dirname + '/packages'));

app.get('/', (req, res) => {
    res.send("Necropolis is a COBOL package manager");
});

app.listen('8080', () => {
    console.log("Listening on port 8080");
});
