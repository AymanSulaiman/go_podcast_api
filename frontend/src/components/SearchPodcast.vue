<template>
    <div>
        <!-- iOS-style Search Bar -->
        <div class="ios-search-container">
            <svg width="16" height="16" xmlns="http://www.w3.org/2000/svg" class="search-icon" @click="clearSearch">
                <circle cx="6" cy="6" r="5" fill="none" stroke="black" stroke-width="2"/>
                <line x1="10" y1="10" x2="15" y2="15" stroke="black" stroke-width="2"/>
            </svg>
            <input v-model="searchTerm" @keyup.enter="searchContent" placeholder="Search for shows or episodes..." class="search-input">
        </div>

        <!-- List of Shows -->
        <div v-if="shows?.length > 0">
            <h2>Shows</h2>
            <ul>
                <li v-for="show in shows" :key="show.collectionId">
                    <h3>{{ show.collectionName }}</h3>
                    <p>By: {{ show.artistName }}</p>
                    <img :src="show.artworkUrl100" alt="Podcast Artwork">
                    <p>Genre: {{ show.primaryGenreName }}</p>
                </li>
            </ul>
        </div>

        <!-- List of Episodes -->
        <div v-if="episodes?.length > 0">
            <h2>Episodes</h2>
            <ul>
                <li v-for="episode in episodes" :key="episode.trackId">
                    <h3>{{ episode.trackName }}</h3>
                    <p>From: {{ episode.collectionName }}</p>
                    <img :src="episode.artworkUrl600" alt="Episode Artwork">
                    <p>Genre: {{ episode.genres[0]?.name || 'N/A' }}</p>
                    <p>Description: {{ episode.description }}</p>
                    <!-- Added a button to play the episode. When clicked, it emits an event to the parent component -->
                    <button @click="(emitPlayEpisodeepisode)">Play</button>
                    <p>Release Date: {{ new Date(episode.releaseDate).toLocaleDateString() }}</p>
                </li>
            </ul>
        </div>
    </div>
</template>

<script>
import axios from 'axios';

export default {

  data() {
    return {
      searchTerm: '',
      shows: [],
      episodes: []
    };
  },

  methods: {
    searchShows() {
      axios.get(`http://localhost:8080/shows?term=${this.searchTerm}`)
        .then(response => {
          this.shows = response.data.results;
          this.$emit('updateSearchResults', { type: 'shows', results: this.shows });
        })
        .catch(error => {
          console.error("There was an error fetching the data:", error);
        });
    },

    searchEpisodes() {
      axios.get(`http://localhost:8080/episodes?term=${this.searchTerm}`)
        .then(response => {
          this.episodes = response.data;
          this.$emit('updateSearchResults', { type: 'episodes', results: this.episodes });
        })
        .catch(error => {
          console.error("There was an error fetching the data:", error);
        });
    },

    searchContent() {
      this.searchShows();
      this.searchEpisodes();
    },

    // Emit the episode to the parent component
    emitPlayEpisode(episode) {
      this.$emit('playEpisode', episode);
    },

  }
}
</script>
