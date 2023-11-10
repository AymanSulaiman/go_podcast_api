<script>
	import { onMount } from 'svelte';
	let searchTerm = "";
	let shows = [];
	let episodes = [];
	let selectedShow = null;
	let selectedEpisode = null;
	const apiBase = "http://localhost:8000/";
  
	onMount(async () => {
	  // Optionally load initial data here
	});
  
	async function searchShows() {
	  try {
		const response = await fetch(`${apiBase}shows?term=${encodeURIComponent(searchTerm)}`);
		if (!response.ok) {
		  throw new Error(`HTTP error! status: ${response.status}`);
		}
		const data = await response.json();
		shows = data.results;
	  } catch (error) {
		console.error("Error fetching shows:", error);
	  }
	}
  
	async function selectShow(show) {
	  try {
		selectedShow = show;
		const response = await fetch(show.feedUrl);
		if (!response.ok) {
		  throw new Error(`HTTP error! status: ${response.status}`);
		}
		const data = await response.text();
		episodes = parsePodcastFeed(data);
	  } catch (error) {
		console.error("Error fetching episodes:", error);
	  }
	}
  
	function parsePodcastFeed(xmlData) {
	  // This function should parse the XML feed and extract the episodes.
	  // You'll need to implement XML parsing here to convert the feed into JSON format.
	  // This is a placeholder function and won't work until implemented.
	  return [];
	}
  
	function selectEpisode(episode) {
	  selectedEpisode = episode;
	  // You would use the episode details to control the audio player.
	}
  </script>
  
  <!-- Search bar -->
  <div class="input-group mb-3">
	<input type="text" class="form-control" placeholder="Search for podcasts" bind:value={searchTerm}>
	<button class="btn btn-outline-secondary" type="button" on:click={searchShows}>Search</button>
  </div>
  
  <!-- Podcast shows display -->
  <div class="horizontal-scroll">
	{#each shows as show}
	  <div class="card" on:click={() => selectShow(show)}>
		<img src={show.artworkUrl100} class="card-img-top" alt={show.collectionName}>
		<div class="card-body">
		  <h5 class="card-title">{show.collectionName}</h5>
		  <p class="card-text">{show.primaryGenreName}</p>
		</div>
	  </div>
	{/each}
  </div>
  
  <!-- Episodes display -->
  {#if selectedShow}
	<div class="horizontal-scroll">
	  {#each episodes as episode}
		<div class="card" on:click={() => selectEpisode(episode)}>
		  <img src={episode.artworkUrl100} class="card-img-top" alt={episode.trackName}>
		  <div class="card-body">
			<h5 class="card-title">{episode.trackName}</h5>
		  </div>
		</div>
	  {/each}
	</div>
  {/if}
  
  <!-- Player at the bottom of the screen -->
  {#if selectedEpisode}
	<div class="fixed-bottom">
	  <audio src={selectedEpisode.episodeUrl} controls autoplay />
	</div>
  {/if}
  
  <style>
	.horizontal-scroll {
	  display: flex;
	  overflow-x: auto;
	}
	.card {
	  min-width: 200px; /* Adjust card width as needed */
	  margin: 5px;
	  cursor: pointer;
	}
	.fixed-bottom {
	  position: fixed;
	  bottom: 0;
	  width: 100%;
	}
  </style>
  