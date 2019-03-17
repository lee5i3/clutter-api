const express = require('express');
const morgan = require('morgan');

const app = express();
const port = 8080;

// log only 4xx and 5xx responses to console
app.use(morgan('dev', {
    skip: function (req, res) {
        return res.statusCode < 400
    }
}));

app.get('/', function (req, res) {
    res.send('Hello World!')
});

app.listen(port, () => console.log(`Listening on port ${port}`))