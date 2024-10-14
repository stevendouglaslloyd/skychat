function LoginButton() {
    const loginButton = document.getElementById("login");
    if (loginButton) {
        loginButton.addEventListener("click", function(event) {
            event.preventDefault(); // Prevent the form from submitting the traditional way

            const username = document.getElementById('username').value;
            const password = document.getElementById('password').value;

            const userObject = {
                username: username,
                password: password
            };

            fetch('/LogonAccount', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(userObject)
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok: ' + response.statusText);
                }
                return response.text(); // or response.json() if expecting JSON response
            })
            .then(data => {
                // Check for incorrect login message
                const incorrectLoginMessage = "Received: Password does not match";
                if (incorrectLoginMessage === data) {
                    document.getElementById('response').innerText = 'Error: ' + data;
                    console.log(data);
                } else {
                    const enterBox = document.querySelector('.enter-box');
                    if (enterBox) {
                        enterBox.remove(); // This will remove the entire section
                    }
                    console.log(data);
                }
            })
            .catch((error) => {
                document.getElementById('response').innerText = 'Error: ' + error.message;
            });
        });
    }
}


function CreateButton() {
    const createButton = document.getElementById("create");
    if (createButton) {
        createButton.addEventListener("click", function(event) {
            event.preventDefault(); // Prevent the form from submitting the traditional way

            const username = document.getElementById('username').value;
            const password = document.getElementById('password').value;

            const userObject = {
                username: username,
                password: password
            };

            fetch('/CreateAccount', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(userObject)
            })

            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok ' + response.statusText);
                }
                return response.text(); // or response.json() if expecting JSON response
            })
            .then(data => {
                // Remove the entire enter box
				const enterBox = document.querySelector('.enter-box');
				var cleanedData = data.replace(/[&{}]/g, '');
				if (enterBox) {
					document.getElementById("response").innerHTML = cleanedData
				}
				
				
				
				console.log(data)
				

                
            })
            .catch((error) => {
                document.getElementById('response').innerText = 'Error: ' + error.message;
            });
        });
    }
}

function SendButton() {
    const sendButton = document.getElementById("send");
    if (sendButton) {
        sendButton.addEventListener("click", function(event) {
            event.preventDefault(); // Prevent the form from submitting the traditional way

            const message = document.getElementById('input').value;
            const time = Date() 
			console.log(message)			
            const MessageBody = {
                message: message,
                time: time
            };

            fetch('/MessageSendTo', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(MessageBody)
            })

            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok ' + response.statusText);
                }
                return response.text(); // or response.json() if expecting JSON response
            })
            .then(data => {
                // Remove the entire enter box

					document.getElementById("messages-box").innerHTML = '<div id="message-entry-sender" class="message-entry-sender"></div>'
					document.getElementById("message-entry-sender").innerHTML = data

				console.log(data)
				

                
            })
            .catch((error) => {
                document.getElementById('response').innerText = 'Error: ' + error.message;
            });
        });
    }
}

function AddButton() {
    const addButton = document.getElementById("Add");
    if (addButton) {
        addButton.addEventListener("click", function(event) {
            event.preventDefault(); // Prevent the form from submitting the traditional way

            const UserName = document.getElementById('InputUserName').value;	
            const PublicKey = document.getElementById('InputPublicKey').value;			
            const ContactJsonStruct = {
				UserName: UserName,
                PublicKey: PublicKey
            };
            

            fetch('/AddContact', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(ContactJsonStruct)
            })

            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok ' + response.statusText);
                }
                return response.text(); // or response.json() if expecting JSON response
            })
            .then(data => {
                // Remove the entire enter box
					//const existingDiv = document.querySelector(".contacts-box"); // Replace with the actual class or id
					//existingDiv.classList.add(".contact-entry");

					document.getElementById("contacts-box").innerHTML += '<div id="contact-entry"></div>'
					document.getElementById("contact-entry").innerHTML += data
					console.log(data);
				

                
            })
            .catch((error) => {
                console.log(error.message);
            });
        });
    }
}
