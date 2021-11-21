/*
* Written by Rajat Banerjee <rajat.banerjee.sjsu.edu>
*/
const fs = require('fs')
const html = fs.readFileSync('index.html','utf-8')

exports.handler = async (event) => {
    const response = {
        statusCode: 200,
        body: html,
        headers: {
            'Content-Type': 'text/html',
        }
    };
    return response;
};
