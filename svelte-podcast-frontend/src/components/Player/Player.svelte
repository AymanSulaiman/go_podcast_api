<script>
	import { onMount } from "svelte";
	import EpisodeModal from "./EpisodeModal.svelte";
	
	// State variables
	let searchTerm = "";
	let shows = [];
	let episodes = [];
	let selectedShow = null;
	let currentAudioUrl = "";
	let audioElement; // Reference to the audio element
	let menuOpen = false;
	let sidebarOpen = false;
	let playlists = [
	  { id: 1, name: 'Playlist 1', episodes: [] },
	  // ... other playlist objects
	];
	
	// Drag and drop state
	let draggedEpisode = null;
	
	// Lifecycle hook for initialization
	onMount(() => {
	  // Add an event listener to the window for spacebar press
	  window.addEventListener("keyup", handleSpacebarPress);
	  return () => {
		// Remove the event listener when the component is destroyed
		window.removeEventListener("keyup", handleSpacebarPress);
	  };
	});
	
	// Search function
	async function searchShowsAndEpisodes() {
	  if (!searchTerm) {
		shows = [];
		episodes = [];
		return;
	  }
	  
	  try {
		const showResponse = await fetch(`http://localhost:8000/shows?term=${searchTerm}`);
		shows = (await showResponse.json()).results;
		
		const episodeResponse = await fetch(`http://localhost:8000/episodes?term=${searchTerm}`);
		episodes = await episodeResponse.json();
	  } catch (error) {
		console.error("Error fetching data:", error);
	  }
	}
	
	// UI Interaction functions
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
	
	function playEpisode(episodeUrl) {
	  currentAudioUrl = episodeUrl;
	  if (audioElement) {
		audioElement.src = currentAudioUrl;
		audioElement.play();
	  }
	}
	
	function handleSpacebarPress(e) {
	  if (e.key === " " && audioElement) {
		if (audioElement.paused) {
		  audioElement.play();
		} else {
		  audioElement.pause();
		}
		e.preventDefault(); // To prevent default behavior of spacebar
	  }
	}
  
	// Playlist functions
	function addToPlaylist(episode) {
	  // For simplicity, we're adding episodes to the first playlist.
	  const targetPlaylist = playlists[0];
	  if (!targetPlaylist.episodes.some(e => e.trackId === episode.trackId)) {
		targetPlaylist.episodes.push(episode);
	  }
	}
  
	function removeFromPlaylist(showId) {
	  const targetPlaylist = playlists[0]; // Assuming we're removing from the first playlist
	  targetPlaylist.episodes = targetPlaylist.episodes.filter(episode => episode.trackId !== showId);
	}
  
	function downloadPlaylist() {
	  const blob = new Blob([JSON.stringify(playlists[0].episodes, null, 2)], { type: 'application/json' });
	  const url = URL.createObjectURL(blob);
	  const a = document.createElement('a');
	  a.href = url;
	  a.download = 'playlist.json';
	  a.click();
	  URL.revokeObjectURL(url);
	}
  
	function loadPlaylist(event) {
	  const fileReader = new FileReader();
	  fileReader.readAsText(event.target.files[0], "UTF-8");
	  fileReader.onload = e => {
		const data = JSON.parse(e.target.result);
		playlists[0].episodes = data; // Assuming the data is an array of episodes for the first playlist
	  };
	}
	
	// Toggle functions
	function toggleMenu() {
	  menuOpen = !menuOpen;
	}
	
	function toggleSidebar() {
	  sidebarOpen = !sidebarOpen;
	}
	
	// Navigation functions
	function goToLogin() {
	  console.log('Login button clicked. Implement login logic.');
	}
  
	// Drag and Drop functions
	function handleDragStart(event, episode) {
	  draggedEpisode = episode;
	  event.dataTransfer.setData('text/plain', episode.trackId); // Set the episode id as the drag data
	}
  
	function handleDragOver(event) {
	  event.preventDefault(); // Necessary for the drop event to fire.
	}
  
	function handleDropOnPlaylist(event, playlistId) {
	  event.preventDefault();
	  const episodeId = event.dataTransfer.getData('text/plain');
	  const episode = episodes.find(ep => ep.trackId.toString() === episodeId);
	  if (episode) {
		const targetPlaylist = playlists.find(p => p.id === playlistId);
		if (targetPlaylist && !targetPlaylist.episodes.some(e => e.trackId === episode.trackId)) {
		  targetPlaylist.episodes.push(episode);
		}
		draggedEpisode = null; // Clear the dragged episode
	  }
	}
  </script>
  
<!-- You'll need to make each episode draggable by adding draggable="true" and the on:dragstart handler -->
<div class="search-results">

	<!-- Episodes Results -->
	{#if episodes.length > 0}
	<h2>Episodes</h2>
	<div class="horizontal-scroll">
		<ul class="episodes-list">
		{#each episodes as episode (episode.trackId)}
		<li draggable="true" on:dragstart={(event) => handleDragStart(event, episode)}>
			<!-- Episode Content -->
		</li>
		{/each}
		</ul>
	</div>
	{/if}
  </div>
  
<!-- Display Playlists -->
	{#each playlists as playlist}
	<div class="playlist" on:dragover={handleDragOver} on:drop={event => handleDropOnPlaylist(event, playlist.id)}>
		<h2>{playlist.name}</h2>
		<ul>
		{#each playlist.episodes as episode}
			<li>{episode.trackName}</li>
		{/each}
		</ul>
	</div>
	{/each}
	

<!-- Search input -->
<input bind:value={searchTerm} on:keyup={handleKeyUp} placeholder="Search for podcasts or episodes..." />

<!-- Toggle Sidebar Button -->
<button on:click={toggleSidebar} class="toggle-sidebar">Toggle Sidebar</button>

<!-- Sidebar for playlists -->
{#if sidebarOpen}
<aside class:sidebar-open={sidebarOpen}>
  <!-- Content for your playlists goes here -->
  <h2>Playlists</h2>
  <!-- Example list of playlists (You would generate this list dynamically based on actual data) -->
  <ul class="playlist-menu">
    <li>Playlist 1</li>
    <li>Playlist 2</li>
    <li>Playlist 3</li>
  </ul>
</aside>
{/if}

<!-- Main Content -->
<div class="search-results">
	<!-- Shows Results -->
	<div class="shows-section">
	  {#if shows.length > 0}
		<h2>Podcasts</h2>
		<div class="horizontal-scroll">
		  <ul class="shows-list">
			{#each shows as show (show.collectionId)}
			  <li class="show-item">
				<div class="show-content">
				  <img class="show-image" src={show.artworkUrl600} alt={show.collectionName} />
				  <div class="show-details">
					<span class="show-title">{show.collectionName}</span>
					<!-- Add more details if needed -->
				  </div>
				</div>
			  </li>
			{/each}
		  </ul>
		</div>
	  {/if}
	</div>
  
<!-- Episodes Results -->
{#if episodes.length > 0}
  <h2>Episodes</h2>
  <div class="episodes-conveyor-container">
    <div class="episodes-conveyor">
      {#each episodes as episode (episode.trackId)}
        <div class="episode-card" draggable="true" on:dragstart={(event) => handleDragStart(event, episode)}>
          <img class="episode-image" src={episode.artworkUrl600} alt={episode.trackName} />
          <h3 class="episode-title">{episode.trackName}</h3>
        </div>
      {/each}
    </div>
  </div>
{/if}




  

  <!-- Playlist Display -->
  {#each playlists as playlist}
    <div class="playlist" on:dragover={handleDragOver} on:drop={(event) => handleDropOnPlaylist(event, playlist.id)}>
      <h2>{playlist.name}</h2>
      <ul>
        {#each playlist.episodes as episode}
          <li>{episode.trackName}</li>
        {/each}
      </ul>
    </div>
  {/each}
</div>

<!-- Selected Show Modal -->
{#if selectedShow}
  <EpisodeModal {selectedShow} onClose={closeEpisodeModal} />
{/if}

<!-- Fixed Audio Player -->
<div id="fixed-audio-player">
  <audio bind:this={audioElement} controls preload="none" src={currentAudioUrl}>
    Your browser does not support the audio element.
  </audio>
</div>

<style>
/* Input Styles */
	input {
	padding: 10px;
	font-size: 16px;
	border: 1px solid #ddd;
	border-radius: 4px;
	width: 100%;
	box-sizing: border-box;
	margin-bottom: 10px;
	}

	/* Toggle Sidebar Button Styles */
	.toggle-sidebar {
	position: fixed;
	top: 20px;
	left: 20px;
	z-index: 11; /* Ensure it's above the sidebar */
	}

	/* Sidebar Styles */
	.sidebar {
	transform: translateX(-100%);
	transition: transform 0.3s ease;
	position: fixed;
	left: 0;
	top: 0;
	width: 250px;
	height: 100%;
	background-color: #fff;
	box-shadow: 3px 0 6px rgba(0,0,0,0.16);
	padding: 20px;
	box-sizing: border-box;
	}

	.sidebar-open {
	transform: translateX(0);
	}

	.playlist-menu {
	list-style-type: none;
	padding: 0;
	margin: 0;
	}

	.playlist-menu li {
	padding: 8px;
	cursor: pointer;
	}

	/* Search Results Styles */
	.search-results {
		display: flex;
		flex-direction: column;
		gap: 20px;
	}

	/* Shows Section Styles */
	.shows-section h2 {
		margin-bottom: 16px;
	}

	.horizontal-scroll {
		overflow-x: auto;
		white-space: nowrap;
		padding-bottom: 16px;
	}

	.shows-list {
		display: flex;
		list-style-type: none;
		padding: 0;
		margin: 0;
	}

	.show-item {
		display: inline-block;
		margin: 8px;
		border: 1px solid #ddd;
		border-radius: 4px;
		overflow: hidden; /* Added to ensure the border-radius clips the content */
	}

	.show-content {
		position: relative;
		width: 100%; /* Adjust as needed */
		text-align: center;
	}

	.show-image {
		width: 100%;
		display: block; /* Ensures there are no gaps under the image */
	}

	.show-details {
		position: absolute;
		bottom: 0;
		left: 0;
		right: 0;
		background: rgba(0, 0, 0, 0.5); /* Semi-transparent overlay */
		color: white;
		padding: 8px;
	}

	.show-title {
		font-weight: bold;
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
	}

	/* Episodes List Styles */
	.episodes-list {
		display: flex;
		list-style-type: none;
		padding: 0;
		margin: 0;
	}

	.episodes-list li {
		display: inline-block;
		margin: 8px;
		border: 1px solid #ddd;
		border-radius: 4px;
		cursor: grab; /* Indicates draggable */
	}

	.episode-content {
		padding: 8px;
	}

	/* Fixed Audio Player Styles */
	#fixed-audio-player {
		position: fixed;
		bottom: 0;
		left: 0;
		width: 100%;
		background-color: rgba(0,0,0,0.8);
		padding: 10px 0;
		z-index: 1000;
	}

	audio {
		width: 100%;
		outline: none;
	}


	.burger {
		position: fixed;
		top: 10px;
		left: 10px;
		z-index: 10;
		background: none;
		border: none;
	}

  .menu {
    position: fixed;
    top: 0;
    left: 0;
    width: 250px; /* adjust as needed */
    height: 100%;
    background-color: white; /* or any color you want */
    box-shadow: 2px 0 5px rgba(0,0,0,0.2);
    transform: translateX(-100%);
    transition: transform 0.3s ease;
    z-index: 9;
  }

  .menu a {
    display: block;
    padding: 10px;
    border-bottom: 1px solid #ddd; /* for a line between items */
    text-decoration: none;
    color: black; /* or any color you want */
  }

  /* When menu is open, slide it into view */
  .menu-open .menu {
    transform: translateX(0);
  }
  .sidebar {
    position: fixed;
    left: 0;
    top: 0;
    width: 250px;
    height: 100%;
    background-color: #fff;
    box-shadow: 3px 0 6px rgba(0,0,0,0.16);
    padding: 20px;
    box-sizing: border-box;
    transform: translateX(-260px);
    transition: transform 0.3s ease;
  }

  /* Show the sidebar when sidebarOpen is true */
  .sidebar-open .sidebar {
    transform: translateX(0);
  }

  /* Styles for the toggle button */
  .toggle-sidebar {
    position: fixed;
    top: 20px;
    left: 20px;
    z-index: 11; /* Ensure it's above the sidebar */
  }

  .sidebar {
  transform: translateX(-100%);
  transition: transform 0.3s ease;
  /* rest of your styles */
}

.sidebar-open {
  transform: translateX(0);
}

/* Styles for the episodes conveyor belt container */
.episodes-conveyor-container {
  overflow-x: scroll; /* This will always show the scrollbar */
  white-space: nowrap;
  padding-bottom: 16px; /* Add padding to give space for the scrollbar */
}

/* Webkit browsers like Chrome, Safari */
.episodes-conveyor-container::-webkit-scrollbar {
  height: 8px; /* Height of the horizontal scrollbar */
}

.episodes-conveyor-container::-webkit-scrollbar-track {
  background: #e0e0e0; /* Color of the scrollbar track */
}

.episodes-conveyor-container::-webkit-scrollbar-thumb {
  background: #888; /* Color of the scrollbar thumb */
  border-radius: 4px; /* Rounded corners on the scrollbar thumb */
}

.episodes-conveyor-container::-webkit-scrollbar-thumb:hover {
  background: #555; /* Color of the scrollbar thumb on hover */
}

/* Firefox */
.episodes-conveyor-container {
  scrollbar-color: #888 #e0e0e0; /* Thumb and track color */
  scrollbar-width: thin; /* Width of the scrollbar */
}

.episodes-conveyor {
  display: flex;
  gap: 16px;
}

.episode-card {
  border: 1px solid #ddd;
  border-radius: 8px;
  overflow: hidden;
  background: #fff;
  cursor: grab;
  transition: transform 0.3s ease;
  text-align: center; /* Center the text */
  display: inline-block; /* Make the card inline for the horizontal scroll */
  width: 200px; /* Fixed width for each card */
}

.episode-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 4px 8px rgba(0,0,0,0.1);
}

.episode-image {
  width: 100%;
  height: auto;
  display: block;
}

.episode-title {
  margin: 8px 0;
  font-size: 1em;
  color: #333;
}

/* Hide scrollbar for aesthetic purposes */
.episodes-conveyor-container::-webkit-scrollbar {
  display: none;
}

.episodes-conveyor-container {
  -ms-overflow-style: none;  /* IE and Edge */
  scrollbar-width: none;  /* Firefox */
}

/* You can add additional styles below */


</style>
