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
