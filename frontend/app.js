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








app.get('/', function(req, res) {
     res.sendFile(curr_dir +'/views/landing.html');
});

app.get('/createBurger', function(req, res) {
	res.sendFile(curr_dir +'/views/createBurger.html')
})

app.get('/secondary_landing', isLoggedIn, function(req, res) {
   
	res.sendFile(curr_dir +'/views/secondary_landing.html')
   
})


/*
app.use(function(req,res,next){
  res.locals.currentUser = req.user;
  localStorage.setItem("username", JSON.stringify(req.user))
  console.log(res.locals.currentUser)
  next();
})
*/



app.post("/signup", function(req, res) {
	User.register(new User({username: req.body.username}), req.body.password, function(err, user) {
         if (err) {
         	console.log(err);
         	return res.sendFile(curr_dir +'/views/landing.html');
         }
         passport.authenticate("local")(req, res, function() {
             res.redirect("/secondary_landing");
         });
	});

})


app.post('/login', passport.authenticate("local", {
	 successRedirect: "/secondary_landing",
	 failureRedirect: "/"
}) , function(req, res){
   
});

app.get('/logout', function(req, res) {
    req.logout();
    res.redirect("/")
})

function isLoggedIn(req, res, next) {
    if (req.isAuthenticated()) {
        return next();
    }
    res.redirect("/")
}


app.get('/starters', function(req, res) {
	res.sendFile(curr_dir +'/views/starters.html')
})

app.get('/burger', function(req, res) {
	res.sendFile(curr_dir +'/views/burger.html')
})

app.get('/shakes', function(req, res) {
	res.sendFile(curr_dir +'/views/shakes.html')
})

app.get('/payment', function(req, res) {
    res.sendFile(curr_dir +'/views/payment.html')
})

app.get('/addEmployee', function(req, res) {
    res.sendFile(curr_dir +'/views/addEmployee.html')
})

app.post('/addEmployee', function(req, res) {
    var Firstname = req.body.firstname;

    var LastName = req.body.lastname;
    var Gender = req.body.gender;
    var Age = parseInt(req.body.age, 10);
    var ID = parseInt(req.body.id, 10);
    var Salary = parseInt(req.body.salary, 10);

  axios.post('http://localhost:5000/employee', {
        FirstName: Firstname,
        LastName: LastName,
        Gender: Gender,
        Age: Age,
        ID: ID,
        Salary: Salary
  })
  .then(function (response) {
    console.log(response);
  })
  .catch(function (error) {
    console.log(error);
  });
              res.redirect("/addEmployee");
})


app.get('/showEmployees', function(req, res) {
  localStorage.setItem("username", JSON.stringify(req.user.username))
    console.log(localStorage.getItem("username").replace(/\"/g, ""))
    res.sendFile(curr_dir +'/views/showEmployees.html')
})


app.get('/searchEmployee', function(req, res) {
    res.sendFile(curr_dir +'/views/searchEmployee.html')
})

app.get('/showSearchEmployee', function(req, res) {
    

    res.sendFile(curr_dir +'/views/showSearchEmployee.html')
})


app.post('/showSearchEmployee', function(req, res) {
  
     res.redirect("/showSearchEmployee");
})

app.get('/deleteEmployee', function(req, res) {

  res.sendFile(curr_dir +'/views/deleteEmployee.html')

})


app.post('/deleteEmployee', function(req, res) {
  var ID = req.body.deleteEmployee_id
 
  axios.delete('http://localhost:5000/employee/delete/'+ID)
    .then(function(response) {
     
    })
     .then(function (response) {
    console.log(response);
  })
  .catch(function (error) {
    console.log(error);
  });
    res.redirect("/deleteEmployee");
})

app.post('/payment', function(req, res) {
    var OrderID = "ABJKLDSJFSDLF";
    var CardHolderName = req.body.CardHolderName;
    var cardNumber = parseInt(req.body.cardNumber, 10);
    var cardType = req.body.cardType;
    var amount = parseInt(req.body.amount, 10);
    //var userID = localStorage.getItem("username")
    
     localStorage.setItem("username", JSON.stringify(req.user.username))
     var userID = localStorage.getItem("username").replace(/\"/g, "")


    //var userID = "Bruce.d"

  axios.post('http://18.205.192.131:80/payment', {
        OrderId: OrderID,
        CardHolderName: CardHolderName,
        CardNumber: cardNumber,
        CardType: cardType,
        UserId: userID,
        Amount: amount
  })
  .then(function (response) {
    console.log(response);
  })
  .catch(function (error) {
    console.log(error);
  });
              res.redirect("/payment");
})

app.get('/paymentHistory', function(req, res) {
     res.sendFile(curr_dir + '/views/paymentHistory.html')
})

app.get('/submitComments', function(req, res) {
    res.sendFile(curr_dir + '/views/submitComments.html')
})



app.listen(4000);
p2pServer.listen();
console.log("Running app at port 4000");

