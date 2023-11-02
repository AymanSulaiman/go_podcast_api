<script>
    import EpisodeModal from './EpisodeModal.svelte';
  
    let shows = [];
    let loading = false;  // initially, we are not loading
    let error = null;
    let searchTerm = '';
    let selectedShow = null;
  
    const fetchShows = async () => {
        loading = true;  // Begin loading state
        error = null;    // Reset any previous errors
      
        try {
            const response = await fetch(`http://localhost:8000/search?term=${searchTerm}`);
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            shows = await response.json();
        } catch (err) {
            error = err.message;
        } finally {
            loading = false;  // End loading state
        }
    };
  
    const fetchEpisodes = async (show) => {
      selectedShow = show; // Set the selected show
      // If needed, you can also fetch episode details here in a similar way to fetchShows
    };
  
    const closeEpisodes = () => {
      selectedShow = null; // Reset the selected show
    };
</script>
  
<input bind:value={searchTerm} on:keydown|preventDefault={event => { if (event.key === 'Enter') fetchShows() }} placeholder="Search for shows" />
<button on:click={fetchShows}>Search</button>
  
{#if loading}
    <p>Loading...</p>
{:else if error}
    <p>Error: {error}</p>
{:else if shows && shows.length > 0}
    {#each shows as show (show.collectionName)}
      <div class="show">
        <button on:click={() => fetchEpisodes(show)} on:keyup={event => event.key === 'Enter' && fetchEpisodes(show)}>
            <img src={show.artworkUrl100} alt="{show.collectionName}" />
        </button>
        
        <h2><a href={show.trackViewUrl} target="_blank">{show.collectionName}</a></h2>
        <p>By {show.artistName}</p>
      </div>
    {/each}
{:else}
    <p>No shows found</p>
{/if}
  
<EpisodeModal show={selectedShow} onClose={closeEpisodes} />
