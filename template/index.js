const fs = require('node:fs');
const path = require('node:path');
const generateHtml = require('./template');

const raw = false
const html = generateHtml(raw)
    

const data = {
    source: html,
    viewport: {
        width: 1024,
        height: 768
    }
}

fs.writeFileSync(path.join(__dirname, 'request.json'), JSON.stringify(data));