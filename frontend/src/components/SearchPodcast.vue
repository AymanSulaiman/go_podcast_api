<template>
    <div>
        <!-- Search Bar -->
        <input v-model="searchTerm" @input="searchContent" placeholder="Search for shows or episodes...">
        
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
                    <audio controls :src="episode.episodeUrl"></audio>
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
    searchContent() {
      this.searchShows();
      this.searchEpisodes();
    },
    searchShows() {
      axios.get(`http://localhost:8080/shows?term=${this.searchTerm}`)
        .then(response => {
          this.shows = response.data.results;
        })
        .catch(error => {
          console.error("There was an error fetching the data:", error);
        });
    },
    searchEpisodes() {
      axios.get(`http://localhost:8080/episodes?term=${this.searchTerm}`)
        .then(response => {
          this.episodes = response.data;
        })
        .catch(error => {
          console.error("There was an error fetching the data:", error);
        });
    }
  }
}
</script>
