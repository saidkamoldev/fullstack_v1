document.addEventListener('DOMContentLoaded', function() {
    document.querySelector('#contact-form-button').addEventListener('click', function(event) {
        event.preventDefault(); // Qayta yuklanishni oldini olish

        var Name = document.querySelector('#contact-form-text').value;
        var Age = document.querySelector('#contact-form-number').value;

        var data = {
            Name: Name,
            Age: Age
        };
        
        var xhr = new XMLHttpRequest();
        xhr.open('POST', 'http://127.0.0.1:8080/products', true);
        xhr.setRequestHeader('Content-Type', 'application/json');
        xhr.onload = function() {
            if (xhr.status >= 200 && xhr.status < 300) {
                console.log('Muvaffaqiyatli:', xhr.responseText);
            } else {
                console.error('Xato yuz berdi:', xhr.status, xhr.statusText);
            }
        };
        xhr.onerror = function () {
            console.error('Sorovni bajarishda xato yuz berdi');
        };

        xhr.send(JSON.stringify(data));
    });



});


document.addEventListener('DOMContentLoaded', function() {
    document.getElementById('load-data').addEventListener('click', function() {
        fetch('http://127.0.0.1:8080/products')
        .then(response => response.json())
        .then(products => {
            const productList = document.getElementById('product-list');
            productList.innerHTML = ''; // Har safar yangidan boshlash uchun oldingi ma'lumotlarni tozalaymiz.
            products.forEach(product => {
                const productDiv = document.createElement('div');
                productDiv.innerHTML = `
                    <h2>${product.Name}</h2>
                    <p>ID: ${product.ID}</p>
                    <p>Created At: ${product.CreatedAt}</p>
                    <p>Updated At: ${product.UpdatedAt}</p>
                    <p>Age: ${product.Age}</p>
                `;
                productList.appendChild(productDiv);
            });
        })
        .catch(error => console.error('Error loading the products:', error));
    });
});
