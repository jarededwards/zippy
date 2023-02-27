
document.addEventListener('DOMContentLoaded', function () {

  document.getElementById('my-form').addEventListener('submit', function (e) {
    e.preventDefault();

    var data = {
      name: document.getElementById('name').value,
    };

    fetch('/api/hello', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    }).then(function (response) {
      return response.json();
    }).then(function (data) {
      document.getElementById('output').innerHTML = data.message;
    });
  });
});
