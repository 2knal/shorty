const serverURL = 'https://sm-all-url.herokuapp.com';

const randomString = () => Math.random().toString(26).substring(2, 9)

document.querySelector('#short').value = `${serverURL}/url/${randomString()}`;

async function submitURL() {
    const long = document.querySelector('#long').value;
    const short = document.querySelector('#short').value;
    
    if (long.length > 0 && short.length > 0) {
        // Submit tso backend
        try {
            let data = await fetch(`${serverURL}/url`, {
                method: 'POST',
                headers: {
                    'content-type': 'application/json'
                },
                body: JSON.stringify({
                    short,
                    long
                })
            });
            const helper = (data) => {
                return data.text().then(text => text ? JSON.parse(text) : {})
            }

            let json = await helper(data)

            if (!json['success']) showSnackbar(json['message'])
        } catch (err) {
            // Raise a message
            showSnackbar('Internal Server Error!')
        }
    } else {
        // Raise a message
        showSnackbar('Please fill all the fields properly!')
    }
}
