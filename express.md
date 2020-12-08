# Express

# Mini WebService printing headers

```

const express = require('express')
const app = express()
 
app.get('/', (req, res) => {
    console.log(JSON.stringify(req.headers));
    res.send('Hello World!');
});
app.get('/plotInstances', (req, res) => {
        console.log(JSON.stringify(req.headers));
        res.send('plotInstances!');
});
 
app.listen(3000, () => console.log('Example app listening on port 3000!'))
```

and run `node index.js`

requires nodejs

IDEA: make it generate the curl command that would result in those headers being received
