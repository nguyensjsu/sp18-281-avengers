const express = require('express');
var cors = require('cors')
const app = express();
app.use(cors())


app.use(express.static('.'));


app.listen(9002,()=>{
    console.log("listening on port 9002");
});


app.get('/',(req,res) => {
    res.sendfile('./burger.html');
} );

app.get('/review',(req,res)=>{
    res.sendfile('./PuristReview.html');
});
module.exports = app;