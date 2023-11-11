import React, { useState, useRef } from 'react';
import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import axios from 'axios';

// Helper function to truncate text
const truncateText = (text, length) => {
  return text.length > length ? text.substring(0, length) + '...' : text;
};

const BASE_URL = 'http://localhost:8000'; // Adjust if your backend is running on a different port

const searchShows = async (term) => {
  try {
    const response = await axios.get(`${BASE_URL}/shows`, {
      params: { term }
    });
    return response.data.results; // Accessing the 'results' array in your shows JSON structure
  } catch (error) {
    console.error('Error fetching shows:', error);
    throw error;
  }
};

const searchEpisodes = async (term) => {
  try {
    const response = await axios.get(`${BASE_URL}/episodes`, {
      params: { term }
    });
    return response.data; // Assuming episodes JSON structure is an array
  } catch (error) {
    console.error('Error fetching episodes:', error);
    throw error;
  }
};

function App() {
  const [searchQuery, setSearchQuery] = useState('');
  const [shows, setShows] = useState([]);
  const [episodes, setEpisodes] = useState([]);
  const [currentEpisodeUrl, setCurrentEpisodeUrl] = useState(''); // State to hold the current episode URL
  const audioPlayer = useRef(null); // Ref for the audio player
  const [currentEpisodeArt, setCurrentEpisodeArt] = useState(''); // State to hold the current episode cover art

  const handleSearchChange = (e) => {
    setSearchQuery(e.target.value);
  };

  const handleSearch = async () => {
    try {
      const fetchedShows = await searchShows(searchQuery);
      const fetchedEpisodes = await searchEpisodes(searchQuery);
      setShows(fetchedShows);
      setEpisodes(fetchedEpisodes);
    } catch (error) {
      console.error('Error during search:', error);
    }
  };

  const playEpisode = (episodeUrl, episodeArt) => { // Update to take episodeArt as parameter
    setCurrentEpisodeUrl(episodeUrl);
    setCurrentEpisodeArt(episodeArt); // Set the artwork when playing an episode
    if (audioPlayer.current) {
      audioPlayer.current.pause();
      audioPlayer.current.load();
      audioPlayer.current.play();
    }
  };

  return (
    <div className="App">
      <div className="search-bar">
        <input
          type="text"
          className="form-control"
          placeholder="Search for podcasts"
          value={searchQuery}
          onChange={handleSearchChange}
        />
        <button onClick={handleSearch} className="btn btn-primary">
          Search
        </button>
      </div>
      {/* Shows Section */}
      <div className="container mt-3">
        <h2>Shows</h2>
        <div className="d-flex flex-nowrap overflow-auto">
          {shows.map((show, index) => (
            <div key={index} className="card m-2" style={{ minWidth: '18rem' }}>
              <img src={show.artworkUrl600} className="card-img-top" alt="Show Artwork" />
              <div className="card-body">
                <h5 className="card-title">{show.collectionName}</h5>
                <p className="card-text">{show.primaryGenreName}</p>
                <a href={show.collectionViewUrl} className="btn btn-primary" target="_blank" rel="noopener noreferrer">View Show</a>
              </div>
            </div>
          ))}
        </div>
      </div>
      {/* Episodes Section */}
      <div className="container mt-3">
        <h2>Episodes</h2>
        <div className="d-flex flex-nowrap overflow-auto">
          {episodes.map((episode, index) => (
            <div key={index} className="card m-2" style={{ minWidth: '18rem' }}>
              <img src={episode.artworkUrl600} className="card-img-top" alt="Episode Artwork" />
              <div className="card-body">
                <h5 className="card-title">{episode.trackName}</h5>
                <p className="card-text">{truncateText(episode.description, 100)}</p>
                <button className="btn btn-primary" onClick={() => playEpisode(episode.episodeUrl, episode.artworkUrl600)}>Play Episode</button>
              </div>
            </div>
          ))}
        </div>
      </div>

      {/* Audio Player */}
      <div style={{ position: 'fixed', bottom: 0, left: 0, width: '100%', backgroundColor: '#fff', display: 'flex', alignItems: 'center', padding: '10px' }}>
        {currentEpisodeArt && (
          <img src={currentEpisodeArt} alt="Cover Art" style={{ height: '75px', marginRight: '10px' }} />
        )}
        <audio ref={audioPlayer} controls style={{ flex: 1 }}>
          <source src={currentEpisodeUrl} type="audio/mpeg" />
          Your browser does not support the audio element.
        </audio>
      </div>
    </div>
  );
}

export default App;