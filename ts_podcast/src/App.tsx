import React, { useState, useEffect } from 'react';
import './App.css';
import logo from './logo.svg';


interface Show {
  wrapperType: string;
  kind: string;
  collectionId: number;
  trackId: number;
  artistName: string;
  collectionName: string;
  trackName: string;
  collectionCensoredName: string;
  trackCensoredName: string;
  collectionViewUrl: string;
  feedUrl: string;
  trackViewUrl: string;
  artworkUrl30: string;
  artworkUrl60: string;
  artworkUrl100: string;
  collectionPrice: number;
  trackPrice: number;
  collectionHdPrice: number;
  releaseDate: string;
  collectionExplicitness: string;
  trackExplicitness: string;
  trackCount: number;
  trackTimeMillis: number;
  country: string;
  currency: string;
  primaryGenreName: string;
  contentAdvisoryRating: string;
  artworkUrl600: string;
  genreIds: string[];
  genres: string[];
}

interface ShowsResponse {
  resultCount: number;
  results: Show[];
}




interface Episode {
  artistIds: number[];
  artworkUrl160: string;
  artworkUrl60: string;
  artworkUrl600: string;
  closedCaptioning: string;
  collectionId: number;
  collectionName: string;
  collectionViewUrl: string;
  contentAdvisoryRating: string;
  country: string;
  description: string;
  episodeContentType: string;
  episodeFileExtension: string;
  episodeGuid: string;
  episodeUrl: string;
  feedUrl: string;
  genres: {
    id: string;
    name: string;
  }[];
  kind: string;
  previewUrl: string;
  releaseDate: string;
  shortDescription: string;
  trackId: number;
  trackName: string;
  trackTimeMillis: number;
  trackViewUrl: string;
  wrapperType: string;
}

const [episodes, setEpisodes] = useState<Episode[]>([]);
const [selectedEpisode, setSelectedEpisode] = useState<Episode | null>(null);


type EpisodesResponse = Episode[];

interface ShowsResponse {
  resultCount: number;
  results: Show[];
}



// Fetch functions
const fetchShows = async (term: string): Promise<ShowsResponse> => {
  const response = await fetch(`http://localhost:8000/shows?term=${encodeURIComponent(term)}`);
  if (!response.ok) {
    throw new Error('Error fetching shows');
  }
  const data: ShowsResponse = await response.json();
  return data;
};


const fetchEpisodes = async (term: string): Promise<EpisodesResponse> =>{
  const response = await fetch(`http://localhost:8000/episodes?term=${encodeURIComponent(term)}`);
  if (!response.ok) {
    throw new Error('Error fetching episodes');
  }
  const data: EpisodesResponse = await response.json();
  return data;
};



const App = () => {
  const [searchTerm, setSearchTerm] = useState(''); // State for the search term
  const [shows, setShows] = useState<Show[]>([]); // State for shows
  const [episodes, setEpisodes] = useState<Episode[]>([]); // State for episodes
  const [selectedEpisode, setSelectedEpisode] = useState<Episode | null>(null); // State for the selected episode

  useEffect(() => {
    const fetchData = async () => {
      if (!searchTerm) {
        setShows([]);
        setEpisodes([]);
        return;
      }

      try {
        const showsData = await fetchShows(searchTerm);
        const episodesData = await fetchEpisodes(searchTerm);

        setShows(showsData.results);
        setEpisodes(episodesData);
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };

    fetchData();
  }, [searchTerm]);

  const handleSearchChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearchTerm(e.target.value);
  };

  return (
    <div>
      <input type="text" value={searchTerm} onChange={handleSearchChange} placeholder="Search for podcasts..." />

      <div>
        <h2>Shows</h2>
        {shows.map(show => (
          <div key={show.collectionId} onClick={() => console.log('Show selected:', show)}>
            {show.collectionName}
            {/* Show more details */}
          </div>
        ))}

        <h2>Episodes</h2>
        {episodes.map(episode => (
          <div key={episode.trackId} onClick={() => setSelectedEpisode(episode)}>
            {episode.trackName}
            {/* Episode more details */}
          </div>
        ))}
      </div>

      {selectedEpisode && (
        <div style={{ position: 'fixed', bottom: 0, left: 0, right: 0, backgroundColor: '#ddd', padding: '10px' }}>
          <h3>Now Playing: {selectedEpisode.trackName}</h3>
          <audio controls src={selectedEpisode.episodeUrl}>
            Your browser does not support the audio element.
          </audio>
        </div>
      )}
    </div>
  );
};

export default App;
