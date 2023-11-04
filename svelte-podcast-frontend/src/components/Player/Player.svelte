<script>
	import { onMount } from "svelte";
	import EpisodeModal from ".EpisodeModal.svelte";

	let searchTerm = "";
	let shows = [];
	let episodes = [];
	let selectedShow = null;
	let currentAudioUrl = "";
	let audioElement; // Reference to the audio element

	async function searchShowsAndEpisodes() {
		if (!searchTerm) {
			shows = [];
			episodes = [];
			return;
		}

		try {
			const showResponse = await fetch(`http://localhost:8000/shows?term=${searchTerm}`);
			shows = await showResponse.json();
			shows = shows.results;

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

	function playEpisode(episodeUrl) {
    currentAudioUrl = episodeUrl;
    if (audioElement) {
        audioElement.src = currentAudioUrl;
        audioElement.play();
    }

	onMount(() => {
		// Add an event listener to the window for spacebar press
		window.addEventListener("keyup", handleSpacebarPress);
		return () => {
			// Clean up the event listener when the component is destroyed
			window.removeEventListener("keyup", handleSpacebarPress);
		};
	});

	function handleSpacebarPress(e) {
		if (e.key === " " && audioElement) {
			if (audioElement.paused) {
				audioElement.play();
			} else {
				audioElement.pause();
			}
			e.preventDefault();  // To prevent default behavior of spacebar (e.g. page scrolling)
		}
	}
}

	
</script>

<input bind:value={searchTerm} on:keyup={handleKeyUp} placeholder="Search for podcasts or episodes..." />

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
						<button 
							class="content" 
							on:click={() => playEpisode(episode.episodeUrl)}
							on:keydown={(e) => e.key === 'Enter' && playEpisode(episode.episodeUrl)}
						>
							<img src={episode.artworkUrl600} alt={episode.trackName} />
							<span class="title">{episode.trackName}</span>
						</button>
					</li>
					{/each}
				</ul>
			</div>
		{/if}
	</div>

	{#if selectedShow}
		<EpisodeModal {selectedShow} onClose={closeEpisodeModal} />
	{/if}

	<div id="fixed-audio-player">
		<audio bind:this={audioElement} controls preload="none" src={currentAudioUrl}>
			Your browser does not support the audio element.
		</audio>
	</div>
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

	/* Search Results Styles */
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

	.content {
		position: relative;
	}

	.overlay {
		position: absolute;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		display: flex;
		align-items: center;
		justify-content: center;
		background-color: rgba(0,0,0,0.5);
		opacity: 0;
		transition: opacity 0.3s;
	}

	.content:hover .overlay {
		opacity: 1;
	}

	.show-button {
		background: none;
		border: none;
		color: inherit;
		padding: 0;
		cursor: pointer;
		outline: inherit;
		text-align: left;
		width: 100%;
	}

	.show-button:hover, .show-button:focus {
		background-color: #f7f7f7;
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
</style>
