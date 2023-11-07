<script>
	import { onMount, onDestroy } from 'svelte';
	import { user } from './stores/userStore';
	import Login from './components/Login/Login.svelte';
	import Player from './components/Player/Player.svelte';
	import logo from './assets/PodCherry_Logo.png';
  
	let isLoggedIn = false;
  
	// Subscribe to user store
	const unsubscribe = user.subscribe(value => {
	  isLoggedIn = value !== null;
	});
  
	onMount(() => {
	  // Initialization logic
	});
  
	onDestroy(() => {
	  unsubscribe();
	});
  
	let showPlayer = true;
  
	function toggleLogin() {
	  showPlayer = !showPlayer;
	}
  
	function showPlayerScreen() {
	  showPlayer = true;
	}
  
	function handleLogoKeydown(event) {
	  if (event.key === 'Enter' || event.key === ' ') {
		showPlayerScreen();
	  }
	}
  </script>
  
  <div class="content-container">
	<button class="logo-container" on:click={showPlayerScreen} on:keydown={handleLogoKeydown}>
	  <img src={logo} alt="PodCherry Logo" class="clickable-logo" />
	</button>
	{#if !isLoggedIn}
	  <button on:click={toggleLogin} class="login-button">Login</button>
	{/if}
  </div>
  
  {#if showPlayer}
	<Player />
  {:else}
	<Login />
  {/if}
  
  



<style>

	.clickable-logo {
		cursor: pointer; /* Indicates that the logo is clickable */
	}

	.logo-container {
		background: none;
		border: none;
		padding: 0;
		cursor: pointer;
	}

	.logo-container:focus {
		outline: none; /* You can style this as needed */
	}

	  /* Style for the login button */
	  .login-button {
		padding: 10px 20px;
		margin-top: 10px; /* Adjust as needed */
		background-color: #007bff; /* Bootstrap primary color */
		color: white;
		border: none;
		border-radius: 5px;
		cursor: pointer;
		font-size: 1em;
	}

	/* Change login button style on hover */
	.login-button:hover {
		background-color: #0056b3;
	}

	.content-container {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}

	.logo-container {
		display: flex;
		justify-content: center; /* Aligns content to the left */
		align-items: center; /* Vertically centers the content in the container */
		padding-top: 20px; /* Padding at the top */
		padding-bottom: 20px; /* Padding at the bottom */
		padding-left: 20px; /* Padding on the left */
	}
  
	/* Styles for the image */
	.logo-container img {
		max-width: 100%; /* Ensures image is responsive and doesn't overflow its container */
		height: auto; /* Maintains the aspect ratio of the image */
		max-height: 50px; /* Example fixed height - adjust as necessary for your layout */
		text-align: center;
	}

	.button-container {
		text-align: center;
		margin-top: 20px; /* Adjust as needed */
	}

	/* Style for the button */
	.button-container button {
		padding: 10px 20px; /* Adjust as needed */
		background-color: #f00; /* Example background color */
		color: #fff; /* Example text color */
		border: none;
		border-radius: 5px; /* Rounded corners */
		cursor: pointer; /* Change cursor on hover */
		font-size: 1em; /* Adjust as needed */
	}

	/* Change button style on hover */
	.button-container button:hover {
		background-color: #c00; /* Darker background on hover */
	}

	/* Mobile devices */
	@media (max-width: 768px) {
		.logo-container img {
			max-height: 30px; /* Smaller size for mobile */
		}
	}

	/* Desktop devices */
	@media (min-width: 1920px) {
		.logo-container img {
			max-height: 140px; /* Larger size for desktop */
		}
	}
</style>


