<template>
  <div id="app">
    <search-podcast></search-podcast>
    <div class="theme-toggle">
      <input type="checkbox" id="theme-switch" @change="toggleTheme" ref="themeSwitch">
      <label for="theme-switch" class="slider"></label>
    </div>
  </div>
</template>

<script>
import SearchPodcast from './components/SearchPodcast.vue';
import { lightTheme, darkTheme } from './themes.js';

export default {
  components: {
    'search-podcast': SearchPodcast
  },
  methods: {
    setTheme(themeName) {
      console.log("Setting theme to:", themeName);
      let theme;
      if (themeName === 'light') {
        theme = lightTheme;
      } else if (themeName === 'dark') {
        theme = darkTheme;
      }

      for (const [key, value] of Object.entries(theme)) {
        document.documentElement.style.setProperty(key, value);
      }

      // Save theme choice to localStorage
      localStorage.setItem('theme', themeName);
    },
    toggleTheme(event) {
      if (event.target.checked) {
        this.setTheme('dark');
      } else {
        this.setTheme('light');
      }
    }
  },
  mounted() {
    const savedTheme = localStorage.getItem('theme');
    if (savedTheme) {
      this.setTheme(savedTheme);
      if(savedTheme === 'dark') {
        this.$refs.themeSwitch.checked = true;
      }
    } else {
      this.setTheme('light');
    }
  }
};
</script>



<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: var(--text);
  margin-top: 60px;
}

:root {
    --color-primary: var(--primary);
    --color-secondary: var(--secondary);
    --color-accent: var(--text);
    --color-neutral: var(--background);
}

body {
  background-color: var(--primary-bg-color);
  color: var(--primary-text-color);
}

#theme-buttons {
  position: sticky;
  top: 0;
  z-index: 10;  /* Ensure buttons are on top */
}

.theme-toggle {
  position: relative;
  width: 60px;
  height: 34px;
}

.theme-toggle input {
  opacity: 0;
  width: 0;
  height: 0;
}

.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #ccc;
  transition: 0.4s;
  border-radius: 34px;
}

.slider:before {
  position: absolute;
  content: "";
  height: 26px;
  width: 26px;
  left: 4px;
  bottom: 4px;
  background-color: white;
  transition: 0.4s;
  border-radius: 50%;
}

input:checked + .slider {
  background-color: #2196F3;
}

input:checked + .slider:before {
  transform: translateX(26px);
}

</style>
