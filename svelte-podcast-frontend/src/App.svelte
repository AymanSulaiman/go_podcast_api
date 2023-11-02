<script>
	import { onMount } from "svelte";
	import EpisodeModal from "./EpisodeModal.svelte";
  
	let searchTerm = "";
	let shows = [];
	let episodes = [];
	let selectedShow = null;
  
	async function searchShowsAndEpisodes() {
	  if (!searchTerm) {
		shows = [];
		episodes = [];
		return;
	  }
  
	  try {
		const showResponse = await fetch(`http://localhost:8000/shows?term=${searchTerm}`);
		shows = await showResponse.json();
		shows = shows.results
  
		const episodeResponse = await fetch(`http://localhost:8000/episodes?term=${searchTerm}`);
		episodes = await episodeResponse.json();
	  } catch (error) {
		console.error("Error fetching data:", error);
	  }
	}
  
	function selectShow(show) {
	  selectedShow = show;
	}
  
	function closeEpisodeModal() {
	  selectedShow = null;
	}
  
	function handleKeyUp(event) {
    if (event.key === 'Enter') {
        searchShowsAndEpisodes();
    }
}
</script>

<!-- <input bind:value={searchTerm} on:input={searchShowsAndEpisodes} on:keyup={handleKeyUp} placeholder="Search for podcasts or episodes..." /> -->
<input bind:value={searchTerm} on:keyup={handleKeyUp} placeholder="Search for podcasts or episodes..." />

<!-- Search Results -->
<div class="search-results">

<!-- Shows Results -->
<div class="shows-section">
	{#if shows.length > 0}
	  <h2>Podcasts</h2>
	  <div class="horizontal-scroll">
		<ul class="shows-list">
		  {#each shows as show (show.collectionId)}
			<li>
			  <button class="show-button" 
				on:click={() => selectShow(show)}
				on:keyup={(e) => e.key === 'Enter' && selectShow(show)}
			  >
				<div class="content">
				  <img src={show.artworkUrl600} alt={show.collectionName} />
				  <span class="title">{show.collectionName}</span>
				</div>
			  </button>
			</li>
		  {/each}
		</ul>
	  </div>
	{/if}
  
	<!-- Episodes Results -->
	{#if episodes.length > 0}
	  <h2>Episodes</h2>
	  <div class="horizontal-scroll">
		<ul class="episodes-list">
		  {#each episodes as episode (episode.trackId)}
			<li>
			  <div class="content">
				<img src={episode.artworkUrl600} alt={episode.trackName} />
				<span class="title">{episode.trackName}</span>
				<a href={episode.trackViewUrl} target="_blank">Play</a>
			  </div>
			</li>
		  {/each}
		</ul>
	  </div>
	{/if}
  
  </div>
  
</div>

{#if selectedShow}
  <EpisodeModal {selectedShow} onClose={closeEpisodeModal} />
{/if}


{#if selectedShow}
  <EpisodeModal {selectedShow} onClose={closeEpisodeModal} />
{/if}

<style>
	input {
		padding: 10px;
		font-size: 16px;
		border: 1px solid #ddd;
		border-radius: 4px;
		width: 100%;
		box-sizing: border-box;
		margin-bottom: 10px;
	}

	.search-results {
		display: flex;
		flex-direction: column;
		gap: 20px; 
	}

	.shows-section, .episodes-section {
		width: 90%;
		margin: 0 auto;
	}

	.horizontal-scroll {
		overflow-x: auto;
		white-space: nowrap;
		padding-bottom: 16px;
	}


	ul {
		display: inline-flex;
		list-style-type: none;
		padding: 0;
		margin: 0;
	}

	li {
		display: inline-block;
		margin: 8px;
		border: 1px solid #ddd;
		border-radius: 4px;
		cursor: pointer;
	}

	img {
		width: 200px;
		height: 200px;
		object-fit: cover;
		display: block;
		margin: 0 auto;
	}

	.title {
		display: block;
		text-align: center;
		padding: 8px 4px;
		font-weight: bold;
		max-width: 220px;
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
	}

	a {
		display: inline-block;
		margin: 8px auto;
		padding: 6px 12px;
		background-color: #007BFF;
		color: #FFF;
		text-decoration: none;
		border-radius: 4px;
		white-space: nowrap;
	}

	a:hover {
		background-color: #0056b3;
	}

	.show-button {
		background: none;
		border: none;
		color: inherit;
		padding: 0;
		cursor: pointer;
		outline: inherit;
		text-align: left;
		width: 100%; /* Adjust according to the image width */
	}

	.show-button:hover, .show-button:focus {
		background-color: #f7f7f7;
	}
</style>
