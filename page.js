function getParameterByName(name) {
    const url = new URL(window.location.href);
    return url.searchParams.get(name);
}

function generateHTMLbyID(id) {
    // Assuming drink.json is accessible via a URL
    fetch('drink.json')
        .then(response => response.json())
        .then(data => {
            const drink = data.find(item => item.id == id);
            if (drink) {
                const container = document.getElementById('container');
                container.innerHTML = `
                    <h1>${drink.name}</h1>
                    <img src="${drink.imageURL}" alt="${drink.name}">
                    <p>Price: ${drink.price.join(', ')}</p>
                    <p>Temperature: ${drink.temperature}</p>
                    <p>Composition: ${drink.composition.join(', ')}</p>
                    <p>Preparation: ${drink.preparation}</p>
                `;
            } else {
                document.getElementById('container').innerHTML = '<p>Drink not found</p>';
            }
        })
        .catch(error => {
            console.error('Error fetching the drink data:', error);
            document.getElementById('container').innerHTML = '<p>Error loading drink data</p>';
        });
}

const id = getParameterByName('id');
generateHTMLbyID(id);
