async function fetchShows() {
    const searchTerm = document.getElementById("searchTerm").value;

    try {
        // Added http:// before localhost
        const response = await fetch(`http://localhost:8080/shows?term=${searchTerm}`);
        if (response.ok) {
            const data = await response.json();
            console.log(data)
            displayResults(data.results);
        } else {
            console.error('Failed to fetch data');
        }
    } catch (error) {
        console.error('Error:', error);
    }
}

function displayResults(shows) {
    const resultsDiv = document.getElementById("results");
    resultsDiv.innerHTML = ''; // clear previous results

    shows.forEach(show => {
        const showDiv = document.createElement('div');
        showDiv.classList.add('show');

        const title = document.createElement('h3');
        title.innerText = show.collectionName;
        showDiv.appendChild(title);

        const image = document.createElement('img');
        image.src = show.artworkUrl100;
        showDiv.appendChild(image);

        const description = document.createElement('p');
        description.innerText = `Artist: ${show.artistName}\nGenre: ${show.primaryGenreName}`;
        showDiv.appendChild(description);

        const link = document.createElement('a');
        link.href = show.collectionViewUrl;
        link.innerText = 'View on Apple Podcasts';
        link.target = '_blank';
        showDiv.appendChild(link);

        resultsDiv.appendChild(showDiv);
    });
}