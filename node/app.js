const express = require('express');
const app = express();

app.listen(3000, () => console.log('express listening on port 3000'));

app.get('/', (req, res) => {
    res.send('<p style="text-align: center; padding-top: 10%;">Aren\'t you a good k8s: <b style="font-size: 2rem;">v1 from node</b></p>');
});