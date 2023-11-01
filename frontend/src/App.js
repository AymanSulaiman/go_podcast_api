// Importing required components and themes
import SearchPodcast from './components/SearchPodcast.vue';
import { lightTheme, darkTheme } from './themes.js';

export default {
  name: 'App',
  components: {
    // Registering the SearchPodcast component
    'search-podcast': SearchPodcast,
  },
  data() {
    return {
      // other properties...
      isPlaying: false,
      playbackPercent: 0
    };
  },
  
  methods: {
    // Function to set the theme based on the given theme name
    setTheme(themeName) {
      console.log("Setting theme to:", themeName);
      let theme;
      // Determining which theme object to use
      if (themeName === 'light') {
        theme = lightTheme;
      } else if (themeName === 'dark') {
        theme = darkTheme;
      }

      // Iterating over theme properties and setting them on the document
      for (const [key, value] of Object.entries(theme)) {
        document.documentElement.style.setProperty(key, value);
      }

      // Saving the theme choice to browser's localStorage for persistence
      localStorage.setItem('theme', themeName);
    },
    // Function to toggle theme based on the checkbox state
    toggleTheme(event) {
      if (event.target.checked) {
        this.setTheme('dark');
      } else {
        this.setTheme('light');
      }
    },
    // Function to clear the search term
    clearSearch() {
      this.searchTerm = '';
    },
    // Function to play the selected episode
    updateSearchResults(payload) {
      if (payload.type === 'shows') {
        this.shows = payload.results;
      } else if (payload.type === 'episodes') {
        this.episodes = payload.results;
      }
    },
    playEpisode(episode) {
      this.currentEpisode = episode;
    },

    togglePlayback() {
      this.isPlaying = !this.isPlaying;

      const audioElem = this.$refs.playerAudio; // Reference to the audio element

      if (this.isPlaying) {
        audioElem.play(); // Play the audio
      } else {
        audioElem.pause(); // Pause the audio
      }
    },

  
    updatePlaybackProgress() {
      if (!this.$refs.playerAudio.duration) return; // Safety check if duration is not available
  
      let currentTime = this.$refs.playerAudio.currentTime;
      let duration = this.$refs.playerAudio.duration;
  
      this.playbackPercent = (currentTime / duration) * 100;
    },
  },
  
  watch: {
    isPlaying(newVal) {
      // If playing, periodically update the playback progress
      if (newVal) {
        this.playbackInterval = setInterval(this.updatePlaybackProgress, 1000); // Update every second
      } else {
        clearInterval(this.playbackInterval); // Clear the interval if paused
      }
    }
  },

  mounted() {
    // When the component is mounted, load saved theme from localStorage
    const savedTheme = localStorage.getItem('theme');
    if (savedTheme) {
      this.setTheme(savedTheme);
      // If saved theme is dark, check the switch
      if(savedTheme === 'dark') {
        this.$refs.themeSwitch.checked = true;
      }
    } else {
      // Default to light theme if no theme is saved in localStorage
      this.setTheme('light');
    }
  }
};
