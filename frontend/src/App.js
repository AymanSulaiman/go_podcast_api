// Importing necessary libraries and styles
import React, { useState, useRef } from 'react';
import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import axios from 'axios';
require('dotenv').config();

// Truncates text to a specified length and adds ellipsis if it's longer than the length
const truncateText = (text, length) => {
  return text.length > length ? text.substring(0, length) + '...' : text;
};

// Base URL for the backend API
const BASE_URL = process.env.backend_url; // Update this if your API is on a different server

// Function to search for shows using a term
const searchShows = async (term) => {
  try {
    const response = await axios.get(`${BASE_URL}/shows`, {
      params: { term }
    });
    // Returns the 'results' array from the API response
    return response.data.results;
  } catch (error) {
    // Logs and throws errors that occur during the API call
    console.error('Error fetching shows:', error);
    throw error;
  }
};

// Function to search for episodes using a term
const searchEpisodes = async (term) => {
  try {
    const response = await axios.get(`${BASE_URL}/episodes`, {
      params: { term }
    });
    // Returns the data assuming the API response structure is an array
    return response.data;
  } catch (error) {
    // Logs and throws errors that occur during the API call
    console.error('Error fetching episodes:', error);
    throw error;
  }
};

// The main App component
function App() {
  // State hooks for various functionalities
  const [searchQuery, setSearchQuery] = useState('');
  const [shows, setShows] = useState([]);
  const [episodes, setEpisodes] = useState([]);
  const [currentEpisodeUrl, setCurrentEpisodeUrl] = useState('');
  const [currentEpisodeArt, setCurrentEpisodeArt] = useState('');
  const [currentShowName, setCurrentShowName] = useState('');
  const [currentShowId, setCurrentShowId] = useState(null);
  const [feedEpisodes, setFeedEpisodes] = useState([]);
  const audioPlayer = useRef(null); // Reference to the audio player DOM element
  const [isShowEpisodesVisible, setIsShowEpisodesVisible] = useState(false);

  // Handles changes to the search input field
  const handleSearchChange = (e) => {
    setSearchQuery(e.target.value);
  };

  // Handles the search operation
  const handleSearch = async () => {
    try {
      // Fetching shows and episodes based on the search query
      const fetchedShows = await searchShows(searchQuery);
      const fetchedEpisodes = await searchEpisodes(searchQuery);
      // Updating the state with the fetched data
      setShows(fetchedShows);
      setEpisodes(fetchedEpisodes);
    } catch (error) {
      // Logs error if search fails
      console.error('Error during search:', error);
    }
  };

  // Plays a selected episode
  const playEpisode = (episodeUrl, episodeArt) => {
    setCurrentEpisodeUrl(episodeUrl);
    setCurrentEpisodeArt(episodeArt);
    // If the audio player is ready, play the episode
    if (audioPlayer.current) {
      audioPlayer.current.pause();
      audioPlayer.current.load();
      audioPlayer.current.play();
    }
  };

  // Views the selected show's episodes by fetching from the feed
  const viewShow = async (feedUrl, showId) => {
    const numericShowId = Number(showId);

    // Toggle visibility off if the same show is selected again
    if (currentShowId === numericShowId) {
      setCurrentShowId(null);
      setCurrentShowName('');
      setFeedEpisodes([]);
      setIsShowEpisodesVisible(false);
    } else {
      // Otherwise, fetch and display the new show's episodes
      try {
        const matchingShow = shows.find(show => show.collectionId === numericShowId);
        if (!matchingShow) {
          throw new Error("Show not found");
        }

        // Update the state with the new show's information
        setCurrentShowName(matchingShow.collectionName);
        setCurrentShowId(numericShowId);
        const response = await axios.post(`${BASE_URL}/fetch-rss`, { url: feedUrl });
        setFeedEpisodes(response.data.Channel.Items);
        setIsShowEpisodesVisible(true);
      } catch (error) {
        console.error('Error fetching feed episodes:', error);
        setCurrentShowId(null);
        setCurrentShowName('');
        setFeedEpisodes([]);
        setIsShowEpisodesVisible(false);
      }
    }
  };


  // Rendering the application UI
  return (
    <div className="App">
      {/* Search bar for entering podcast search terms */}
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

      {/* List of shows returned from the search */}
      <div className="container mt-3">
        <h2>Shows</h2>
        <div className="d-flex flex-nowrap overflow-auto">
          {shows.map((show, index) => (
            <div key={index} className="card m-2" style={{ minWidth: '18rem' }}>
              <img src={show.artworkUrl600} className="card-img-top" alt="Show Artwork" />
              <div className="card-body">
                <h5 className="card-title">{show.collectionName}</h5>
                <p className="card-text">{show.primaryGenreName}</p>
                <button onClick={() => viewShow(show.feedUrl, show.collectionId.toString())} className="btn btn-primary">
                  View Show
                </button>
              </div>
            </div>
          ))}
        </div>
      </div>

      {/* Conditional rendering of episodes from the selected show's feed */}
      {isShowEpisodesVisible && currentShowId && feedEpisodes.length > 0 && (
        <div className="container mt-3">
          <h2>Episodes from {currentShowName}</h2>
          <div className="d-flex flex-column">
            {feedEpisodes.map((episode, index) => (
              <div key={index} className="mb-2">
                <h5>{episode.Title}</h5>
                <p>{truncateText(episode.Description, 100)}</p>
                {/* Additional episode details can go here */}
              </div>
            ))}
          </div>
        </div>
      )}


      {/* List of episodes returned from the search */}
      <div className="container mt-3">
        <h2>Episodes</h2>
        <div className="d-flex flex-nowrap overflow-auto">
          {episodes.map((episode, index) => (
            <div key={index} className="card m-2" style={{ minWidth: '18rem' }}>
              <img src={episode.artworkUrl600} className="card-img-top" alt="Episode Artwork" />
              <div className="card-body">
                <h5 className="card-title">{episode.trackName}</h5>
                <p className="card-text">{truncateText(episode.description, 100)}</p>
                <button className="btn btn-primary" onClick={() => playEpisode(episode.episodeUrl, episode.artworkUrl600)}>
                  Play Episode
                </button>
              </div>
            </div>
          ))}
        </div>
      </div>

      {/* Fixed audio player at the bottom of the screen */}
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
