// Adds item to local storage variable that user is adding to cart.
function addtocart(name, rate) {

    var isChanged = false;

/*    var order = JSON.parse(localStorage.getItem('order'));
    console.log("[DEBUG 01]: ", order)*/

    var item = {
        "name": name,
        "count": 1,
        "rate": rate,
    }

    var order = JSON.parse(localStorage.getItem('order'));
    console.log(order)

    if (order.items.length === 0) {
        order.items.push(item)
    } else {
        // Traverse through item and increase count if item already exist or add another item if it doesn't exist.
        for(var i = 0; i < order.items.length; i++) {
            if (order.items[i].name === name) {
                console.log("inside for");
                order.items[i].count++;
                isChanged = true;
                break;
            }
        }
        if (!isChanged) {
            order.items.push(item)
            // isChanged = false
        }
    }
    localStorage.setItem("order", JSON.stringify(order));

    // console.setItem("[DEBUG]")

    // console.log("Length of items: ", order.items.length)
    console.log("========================")

    for (var k = 0; k < order.items.length; k++) {
        console.log("Name: ", order.items[k].name)
        console.log("Count: ", order.items[k].count)
        
    }


}

// Function called when user clicks on Done after selecting all the items he/she wants to order.
// Also opens another page of cart.
function doneOrder() {

    var order = JSON.parse(localStorage.getItem('order'));

    console.log("[DEBUG]: ", typeof(order))
    console.log("[DEBUG]: ", order)


    axios.post('http://localhost:3000/order', order)
            .then(function (response){
                console.log(response);
                console.log("DEBUG]: Return response", response.data);
                localStorage.setItem("orderid", response.data.id)
                console.log(response.data.id)
                location.href = "cart.html";
            })
            .catch(function(error) {
                console.log(error);
            });
}

/ Populate cart for current order and current user.
function populateCart() {
    console.log("[DEBUG CART]: Populate cart called.")

    var orderId = localStorage.getItem("orderid")

    console.log("[DEBUG CART]: Order ID is: ", orderId);

    axios.get('http://localhost:3000/view/' + orderId)
        .then(function (response) {
            // console.log(response.data.id);
            // console.log("[CART DEBUG] Fetched from database: ", response.data)
            // alert(response.data);

            // console.log("[DEBUG]: ", response.data.items.length)

            var result=[];
            for(var i=0;i<response.data.items.length;i++){
                console.log("[DEBUG LOOP]: Inside for loop.")
                result.push(response.data.items[i]);
                // console.log("[DEBUG DATA STORE]: ", result[i])
            }
            console.log("[DEBUG]: Length of items: ", result.length);

            // Create table structure
            var table = '<style>' +
                'table {' +
                    'font-family: arial, sans-serif;' +
                    'border-collapse: collapse;' +
                    'width: 70%;text-align:center;' +
                    'border-radius:2px}' +
                'td,th {' +
                    'border: 1px solid #dddddd;' +
                    'text-align: center;' +
                    'padding: 8px;}' +
                'th {' +
                    'background-color:#e8c592;}' +
                'td {' +
                    'background-color:#ffffcc;}' +
                '</style>' +
                '<table align="center" cellpadding="2" style="text-align:center;font-family: arial,sans-serif;border-collapse: collapse>"';

            table +='<tr>' +
                    '<th>ITEM NAME</th><th>COUNT</th><th>RATE</th><th>AMOUNT</th>' +
                '</tr>';


            for(var i=0;i<result.length;i++){
                /* console.log("in loop");*/
                console.log("[DEBUG]: Amount: ", result[i].amount);
                table += '<em></em><tr><td>' +
                    result[i].name +
                    '</td><td>' + result[i].count + '</td><td>' + result[i].rate + '</td><td>' + result[i].amount + '</td>' +
                    '</tr></em>';
                //console.log(response.data[i]);
            }
            table += '</table>';
            document.getElementById('itemlist').innerHTML = table;
        })
        .catch(function (error) {
            console.log(error);
        });
}
