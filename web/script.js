var findOrder = document.getElementById("find_order");
findOrder.addEventListener("click", function(){

    var order_uid = document.getElementById('text_find_order').value;

    $.ajax({
        type: "GET",

        url: "/order",
        data: {
            "order_uid": order_uid,
        },
        success: function (data) {
            document.getElementById('result_order').textContent = JSON.stringify(data, null, 2);
            console.log(data);
        },
        failure: function (data) {
            console.log("failure");
            console.log(data);
        },
    });
})