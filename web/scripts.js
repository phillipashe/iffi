function handleFileSelect(event) {
  var file = event.target.files[0];
  var reader = new FileReader();
  
  reader.onload = function(e) {
    var previewImage = document.getElementById('preview');
    previewImage.src = e.target.result;
  };
  
  reader.readAsDataURL(file);
}

function submitProto() {
  var request = new MyProtoMessage(); // Create an instance of your proto message
  request.setField1("Hello");
  request.setField2("World");

  var client = new MyServiceClient('https://your-grpc-server.com');
  var metadata = {}; // Optional metadata for the gRPC call

  client.myMethod(request, metadata, function(err, response) {
    if (err) {
      console.error('Error:', err.message);
    } else {
      console.log('Response:', response.toObject());
    }
  });
}