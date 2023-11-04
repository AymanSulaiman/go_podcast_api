<script>
    export let show = null; // The selected show
    export let onClose = () => {}; // A function to close the modal
  
    let episodes = [];
  
    // When the show changes, fetch its episodes.
    $: if (show) {
      // Fetch episodes related to the show
      // Note: This is a mockup and assumes a specific podcast for simplicity.
      // In a real-world scenario, use an actual API call to fetch episodes.
      if (show.collectionName === 'Danica Patrick Pretty Intense Podcast') {
        episodes = [
          {
            // ... The provided episode details ...
          }
        ];
      }
    }
  </script>
  
  {#if show}
  <div class="modal">
    <button on:click={onClose}>Close</button>
    <h2>Episodes of {show.collectionName}</h2>
    {#each episodes as episode}
      <div class="episode">
        <img src={episode.artworkUrl160} alt="{episode.trackName}" />
        <p>{episode.trackName}</p>
        <a href={episode.trackViewUrl} target="_blank">Listen</a>
      </div>
    {/each}
  </div>
  {/if}
  
  <style>
    /* Modal styling */
    .modal {
      position: fixed;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      width: 80%;
      max-width: 600px;
      background-color: #fff;
      padding: 20px;
      border-radius: 5px;
      z-index: 1000;
      box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
      overflow-y: auto;
    }
  
    .episode {
      display: flex;
      align-items: center;
      margin-bottom: 15px;
    }
  
    .episode img {
      width: 60px;
      height: 60px;
      margin-right: 15px;
    }
  
    .episode p {
      flex-grow: 1;
    }
  </style>
  