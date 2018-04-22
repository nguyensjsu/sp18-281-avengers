var express = require('express')
var app = express();
var bodyParser = require('body-parser');
var path = require('path')
var cookieParser = require('cookie-parser')
var exphbs = require('express-handlebars')
var expressValidator = require('express-validator')
var flash = require('connect-flash');
var session = require('express-session')
var passport = require('passport')
var LocalStrategy = require('passport-local'),Strategy;
var passportLocalMongoose = require('passport-local-mongoose')
var mongo = require('mongodb')
var mongoose = require('mongoose');
var User = require("./models/user");
var axios = require("axios");
var LocalStorage = require('node-localstorage').LocalStorage,
localStorage = new LocalStorage('./scratch');


//Blockchain dependencies
const Blockchain = require('./blockchain');
const P2pServer = require('./p2p-server')


const bc = new Blockchain();
const p2pServer = new P2pServer(bc);
app.use(bodyParser.json())
app.use(bodyParser.json())



/*

axios.get('http://localhost:3000/employees')
    .then(function(response) {
        console.log(response.data[0].firstname)
    });
*/

mongoose.connect('mongodb://localhost/loginapp');

app.use(require("express-session")({
    secret: "A word",
    resave: false,
    saveUninitialized: false
}));
app.use(passport.initialize());
app.use(passport.session());
passport.use(new LocalStrategy(User.authenticate()));
passport.serializeUser(User.serializeUser())
passport.deserializeUser(User.deserializeUser())

var db = mongoose.connection;


var curr_dir = process.cwd()
app.use(express.static("/"));


app.use("/app.js", express.static(__dirname + '/app.js'));

app.use(express.static(path.join(__dirname, 'public')));
app.use(express.static(path.join(__dirname, 'views')));
app.use(bodyParser.urlencoded({limit: '50mb', extended: true}));
app.use(bodyParser.json({limit: '50mb'}));


