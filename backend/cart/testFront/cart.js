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
