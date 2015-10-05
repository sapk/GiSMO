console.log("Ready !");

var Users = Backbone.Collection.extend({
  url: '/api/user'
});


//Movies/Series (Folder) and Video,Episode 
var Items = Backbone.Collection.extend({
  url: '/api/item'
});


var user = {};

_.extend(user, Backbone.Events);

user.on("login", function(pseudo) {
  alert("User logged : " + pseudo);
});

//user.trigger("login", "an event");

$.get("/api/user/_current",function(u){
    console.log(u);
            if(u.isLogged){
                _.extend(user, u);
                user.trigger("login", u.username);
            }else{
                login();
            }
},"json");

function login(){
        //TODO ask for login better
        var username = prompt("Please enter your username", "");
        var password = prompt("Please enter your password", "");
        
        $.post("/api/user/_login",{username:username,password:password},function(u){
            if(u.isLogged){
                _.extend(user, u);
                user.trigger("login", u.username);
            }else{
                login();
            }
        },"json");

}