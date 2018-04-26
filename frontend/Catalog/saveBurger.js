

function deleteBurger()
{
    //document.getElementById('save').style.visibility='hidden';

    var title = document.getElementById("title").value;
    var price = document.getElementById("price").value;
    var description = document.getElementById("description").value;
    var code = document.getElementById("code").value; 

    // var config ={
    //     headers: { 
    //         "test": "abc",
    //         "Content-Type": "application/json",
    //         "Access-Control-Request-Headers": "X-Requested-With",
    //         "Access-Control-Request-Method":"POST"
    //     }
    // }

    var burgerData ={
        Title : title,
        Price : price,
        Description : description,
        Code : code
    }
    console.log(burgerData); 
    axios.delete('http://127.0.0.1:3305/item',{ data: { Title:title,Price:price,Description:description,Code:code }})
      .then(function (response) {
          console.log(response);
          alert(response.data.Item)
          location.href = 'burger.html'
        })
        .catch(function (error) {
          console.log(error);
        });
}


function fetchBurger(){

    axios.get('http://127.0.0.1:3305/displayitem',{ "Content-Type": "application/json"})
    .then(function (response) {
        //console.log(response)
        // console.log(response.data[1]);
        // console.log(response.data[3]);
        // console.log(response.data[5]);
        
        //console.log("in fetch burger");
        //console.log("length",response.data.length);
        //result.push(response.data[1]);
        var result=[];
        for(var i=1;i<response.data.length;i++){
           
            result.push(response.data[i]);
            i++;
        //     console.log("value of fetch: ",response.data[i+1])
        //    console.log(finallist[i]);
           // i++;  
        }
        // console.log(result.length);
        for(var i=0;i<result.length;i++){
           // console.log("in loop");
           console.log(result[i]);
           //console.log(response.data[i]);
        }
        // var s='<span>'+result[3]+'</span></br> <span>'+response.data[1]+'</span></br><span>'+response.data[2]+'</span><button type="button" class="btn btn-success add-cart" id="bacon">Add to Cart</button></label>'
        //     document.getElementById("result").innerHTML=s
        var s = '<style>table {font-family: arial, sans-serif;border-collapse: collapse;width: 70%;text-align:center;border-radius:2px}td,th {border: 1px solid #dddddd;text-align: center;padding: 8px;}th {background-color:#e8c592;}</style><table align="center" cellpadding="2" style="text-align:center;font-family: arial,sans-serif;border-collapse: collapse"';
        s +='<tr></tr>';
        for(var i=0;i<result.length;i++)
        {
            s +='<tr class="tdrow">';
            s +='<td  style="border: 1px solid #dddddd;"><label class="container"><div style="float:left"><span>Code:'+" "+result[i]+'</span></br> <span>'+result[++i]+'</span></br><span>'+result[++i]+'</span></br><span>'+result[++i]+'</span></div><div style="float:right">&nbsp;&nbsp;<button type="button" class="btn btn-success add-cart" id="bacon">Add to Cart</button>&nbsp;&nbsp;<button type="button" class="btn btn-primary moreInfo" id="moreInfo" style="float:right" onclick="review()">Review</button></div><div style="float:right"><input type="text" id="itemcount" style="width:40px;color:black"></div></label></td>';
            s +='</tr>';
        }
        s +='</table>'
        document.getElementById('result').innerHTML=s;
     
      })
      .catch(function (error) {
        console.log(error);
      });
}


function createBurger(){  
    //document.getElementById("delete").style.visibility='hidden';
    var randomCode=Math.floor(Math.random()*1000);
    var code = randomCode+"";
    var title = document.getElementById("title").value;
    var price = document.getElementById("price").value;
    var description = document.getElementById("description").value;
    var productCode=Math.random().toString(36).substring(2,5);
    //console.log("random String is",ranstring);
    // document.getElementById("code").value=code;
    // var productCode=document.getElementById("code").value;
    // console.log("Product code",productCode);
    var burgerData ={
        Title : title,
        Price : price,
        Description : description,
        Code : productCode
    } 
    axios.post('http://localhost:3305/createitem',burgerData)
    .then(function (response) {
        console.log(response.data.Item);
        alert(response.data.Item);
        location.href = 'burger.html'
      })
      .catch(function (error) {
        console.log(error);
      });

}

function updatePrice(){
    var title = document.getElementById("title").value;
    var price = document.getElementById("price").value;
    var code = document.getElementById("code").value;
    var burgerData ={
        Title : title,
        Price : price,
        Description : "description",
        Code : code
    }
      axios.put('http://localhost:3305/item',burgerData)
      .then(function (response) {
          console.log(response);
          alert(response.data.Item);
          location.href = 'burger.html'
        })
        .catch(function (error) {
          console.log(error);
        });

}


function moreinfo(name, description) {

    var review = {
        itemname: name,
        description: description,    
    }

    localstorage.setItem("reviewname", review)

    location.href = "PuristReview.html"
    var s='<span>'+localstorage.getItem("reviewname")+'</span>'
    document.getElementById("result").innerHTML=s;
}