const serverURL = 'https://sm-all-url.herokuapp.com';

async function fetchStats() {
    try {
        let data = await fetch(`${serverURL}/url`);
        data = await data.json();
        const urls = data['data']
        let table = document.querySelector('#add-here')
        
        // Add to table
        urls.forEach(url => {
            table.innerHTML += `
                <tr>
                    <td>${url.short}</td>
                    <td>${url.long}</td>
                    <td>${url.count}</td>
                <tr/>
            `
        });
        
        if (!data['success']) showSnackbar(json['message'])
    } catch (err) {
        // Raise a message
        showSnackbar('Internal Server Error!')
    }
}

// Initializing table
fetchStats()
