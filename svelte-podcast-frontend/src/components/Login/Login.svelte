<!-- src/components/Login/Login.svelte -->
<script>
  import { GoogleAuthProvider, signInWithPopup, OAuthProvider } from "firebase/auth";
  import { signInWithEmailAndPassword, createUserWithEmailAndPassword } from "firebase/auth";
  import { auth } from '../../firebase.js'; // Adjust the relative path as necessary
  import { user } from '../../stores/userStore';
  import { navigate } from 'svelte-routing';

  // Define the provider for Google authentication
  const googleProvider = new GoogleAuthProvider();

  // Define the provider for Apple authentication
  const appleProvider = new OAuthProvider('apple.com');

  let isSignUp = false; // Flag to toggle between login and sign up
  let email = '';
  let password = '';
  let confirmPassword = ''; // Only used for sign up to confirm password
  let errorMessage = ''; // To display any error message from Firebase

  function toggleForm() {
    isSignUp = !isSignUp;
    errorMessage = ''; // Clear error message when toggling
  }

  async function handleLogin() {
    try {
      const response = await signInWithEmailAndPassword(auth, email, password);
      user.set(response.user);
    } catch (error) {
      console.error('Login failed:', error);
      errorMessage = error.message;
    }
  }

  async function handleSignUp() {
    if (password !== confirmPassword) {
      errorMessage = 'Passwords do not match!';
      return;
    }
    try {
      const response = await createUserWithEmailAndPassword(auth, email, password);
      user.set(response.user);
    } catch (error) {
      console.error('Sign up failed:', error);
      errorMessage = error.message;
    }
  }

  async function signInWithGoogle() {
    try {
      const response = await signInWithPopup(auth, googleProvider);
      user.set(response.user);
    } catch (error) {
      console.error('Google sign-in failed:', error);
      errorMessage = error.message;
    }
  }

  // Function to handle Apple sign-in
  async function signInWithApple() {
    try {
      const response = await signInWithPopup(auth, appleProvider);
      user.set(response.user);
    } catch (error) {
      console.error('Apple sign-in failed:', error);
      errorMessage = error.message;
    }
  }

</script>
<link rel="stylesheet" href="https://use.fontawesome.com/releases/v5.8.1/css/all.css">
  
<div>
  <!-- Display error message if there is one -->
  {#if errorMessage}
    <p class="error">{errorMessage}</p>
  {/if}

  {#if isSignUp}
    <h2>Sign Up</h2>
    <input type="email" bind:value={email} placeholder="Email" />
    <input type="password" bind:value={password} placeholder="Password" />
    <input type="password" bind:value={confirmPassword} placeholder="Confirm Password" />
    <button on:click={handleSignUp}>Sign Up</button>

    <button on:click={signInWithGoogle} class="social-signup">
      <i class="fab fa-google"></i> Sign Up with Google
    </button>
    <button on:click={signInWithApple} class="social-signup">
      <i class="fab fa-apple"></i> Sign Up with Apple
    </button>
  {:else}
    <h2>Sign In</h2>
    <input type="email" bind:value={email} placeholder="Email" />
    <input type="password" bind:value={password} placeholder="Password" />
    <button on:click={handleLogin}>Sign In</button>
  {/if}

  {#if !isSignUp}
  <button on:click={signInWithGoogle}>
    <i class="fab fa-google"></i> Sign In with Google
  </button>
  <button on:click={signInWithApple}>
    <i class="fab fa-apple"></i> Sign In with Apple
  </button>
  
  {/if}

  <button on:click={toggleForm}>
    {#if isSignUp}
      Already have an account? Sign In
    {:else}
      Need an account? Sign Up
    {/if}
  </button>
</div>
  
<style>
  div {
    display: flex;
    flex-direction: column;
    justify-content: flex-start; /* Start items at the beginning of the container */
    align-items: center;
    min-height: 50vh; /* Reduce the minimum height */
    padding: 20px;
  }

  input, button {
    margin: 10px 0;
  }

  .error {
    color: red;
  }

  button {
  display: flex;
  align-items: center;
  justify-content: center; /* Add this line to center align children horizontally */
  width: 15%; /* If you want the button to stretch to full container width */
  /* ... other styles ... */
}

  button i {
    margin-right: 8px; /* This will keep the space between the icon and the text */
    /* Since you've already set a fixed width and height for img, you might want to do the same for icons */
    width: 24px; /* Or whatever size you prefer */
    height: 24px; /* Or whatever size you prefer */
  }

  /* If you want to ensure that the text and icon are always in the center regardless of button width: */
  button i, button span {
    flex: none; /* This prevents the children from stretching */
  }

  .social-signup {
    width: 15%; /* Change this to match the style of the other buttons */
    margin-top: 10px;
    /* Add other styles if necessary */
  }
</style>
