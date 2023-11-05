<!-- src/components/Login/Login.svelte -->
<script>
    import { GoogleAuthProvider, signInWithPopup, OAuthProvider } from "firebase/auth";
  import { signInWithEmailAndPassword, createUserWithEmailAndPassword } from "firebase/auth";
  import { auth } from '../../firebase.js'; // Adjust the relative path as necessary
  import { user } from '../../stores/userStore';

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
  
  <div>
    <button on:click={toggleForm}>
      {#if isSignUp}
        Already have an account? Sign In
      {:else}
        Need an account? Sign Up
      {/if}
    </button>
    
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
    {:else}
      <h2>Sign In</h2>
      <input type="email" bind:value={email} placeholder="Email" />
      <input type="password" bind:value={password} placeholder="Password" />
      <button on:click={handleLogin}>Sign In</button>
    {/if}

    {#if !isSignUp}
      <button on:click={signInWithGoogle}>Sign In with Google</button>
      <button on:click={signInWithApple}>Sign In with Apple</button>
    {/if}
  </div>
  
  <style>
    .error {
      color: red;
    }
    /* Add more CSS if needed */
  </style>
  