<script>
  import { onMount } from 'svelte';
  
  let shows = [];
  let loading = true;
  let error = null;
  let searchTerm = '';

  const fetchShows = async () => {
    loading = true;
    error = null;
    try {
      const response = await fetch(`http://localhost:8000/shows?term=${searchTerm}`);
      if (!response.ok) {
        throw new Error('Server responded with ' + response.status);
      }
      const data = await response.json();
      shows = data.results;
    } catch (err) {
      error = err.message;
    } finally {
      loading = false;
    }
  };

  const handleKeyup = (event) => {
    if (event.key === 'Enter') {
      fetchShows();
    }
  };

  onMount(fetchShows);
</script>

<input bind:value={searchTerm} on:keyup={handleKeyup} placeholder="Search for shows" />


{#if loading}
  <p>Loading...</p>
{:else if error}
  <p>Error: {error}</p>
{:else if shows && shows.length > 0}
  {#each shows as {collectionName, artistName, artworkUrl600, trackViewUrl} (collectionName)}
    <div class="show">
      <img src={artworkUrl600} alt="{collectionName}" />
      <h2><a href={trackViewUrl} target="_blank">{collectionName}</a></h2>
      <p>By {artistName}</p>
    </div>
  {/each}
{:else}
  <p>No shows found</p>
{/if}

<style>
  input, button {
    margin-bottom: 1rem;
  }

  .show {
    border: 1px solid #ccc;
    padding: 1rem;
    margin-bottom: 1rem;
    display: flex;
    flex-direction: column;
    align-items: center;
    max-width: 250px; /* Increased the max-width to better accommodate the larger images and text */
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    overflow: hidden;
    background-color: #fff;
  }

  img {
    max-width: 100%;
    height: auto;
    margin-bottom: 1rem;
  }

  h2 {
    margin: 0 0 0.5rem 0;
    font-size: 1.25rem;
    text-align: center;
    font-weight: bold;
  }

  p {
    margin: 0;
    color: #555;
    text-align: center;
  }
</style>
