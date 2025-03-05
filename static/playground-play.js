import { Simulator, upgradeOldState } from './playground/simulator.js';
import { Editor } from './playground/resources/editor.bundle.min.js';
import { boards } from './playground/boards.js';

// This file is for the playground at the /play URL.

const PLAYGROUND_API = location.hostname == 'localhost' ? 'http://localhost:8080/api' : 'https://playground-bttoqog3vq-uc.a.run.app/api';

const defaultBoardName = 'console';

let simulator = null;
let editor = null;

// Regexp for a URL like /s/abcdef
let shareURLRegexp = RegExp('/s/([a-zA-Z0-9]*)$');

// updateBoards updates the dropdown menu. This must be done on load or after
// selecting a different board.
function updateBoards(state) {
  // Try to figure out the board name based on the saved state.
  let activeName = '';
  if (state) {
    let matches = RegExp('parts/([a-z0-9-]+).json$').exec(state.parts[0].location);
    if (matches) {
      activeName = matches[1];
      if (activeName === 'console' && state.compiler === 'go') {
        activeName = 'console-go';
      }
    }
    let button = document.querySelector('#target > button');
    if (activeName in boards) {
      button.textContent = boards[activeName].humanName;
    } else {
      // For some reason we couldn't find the board (should only happen with a
      // malformed editor state).
      button.textContent = '...';
    }
  }

  // Replace the dropdown with a new dropdown.
  let dropdown = document.querySelector('#target > .dropdown-menu');
  dropdown.innerHTML = '';
  for (let name in boards) {
    let item = document.createElement('a');
    item.textContent = boards[name].humanName;
    item.classList.add('dropdown-item');
    if (name === activeName) {
      item.classList.add('active');
    }
    item.setAttribute('href', '');
    item.dataset.name = name;
    dropdown.appendChild(item);
    item.addEventListener('click', (e) => {
      e.preventDefault();
      clearShareURL();
      setBoard(item.dataset.name);
    });
  }
}

// setBoard reloads the editor and simulator to use the given board.
async function setBoard(name) {
  // Create a new clean state object for this board.
  let board = boards[name];
  let state = {
    code: board.code,
    compiler: board.compiler,
    parts: [
      {
        id: 'main',
        location: board.location,
        x: 0,
        y: 0,
      },
    ],
    wires: [],
  };

  // Update the editor/simulator.
  await setState(state);

  // Load the same board after a reload.
  localStorage.tinygo_playground_boardName = name;
  delete localStorage.tinygo_playground_state;
}

// Reset the editor and simulator to use the given state.
async function setState(state) {
  upgradeOldState(state);
  updateBoards(state);
  editor.setText(state.code);

  // Change to the new project state.
  await simulator.setState(state);
}

// Load the editor state on page load.
async function loadSavedState() {
  // Check whether this is a share URL, and if so load the editor state from the
  // API.
  let matches = shareURLRegexp.exec(document.location.pathname);
  if (matches) {
    let id = matches[1];
    let response = await fetch(`${PLAYGROUND_API}/share?id=${id}`);
    if (response.status != 200) {
      // Can't load the shared code.
      // Users can still select a different sample from the dropdown. This is
      // not very obvious, so might need a more explicit message.
      simulator.terminal.appendError(`failed to fetch shared code: ${response.status} ${response.statusText}`);
      return;
    }
    let state = await response.json();
    setState(state.data);
    return;
  }

  // Check whether the editor saved a previous state.
  if (localStorage.tinygo_playground_state) {
    setState(JSON.parse(localStorage.tinygo_playground_state));
    return;
  }

  // Load the default board.
  let board = localStorage.tinygo_playground_boardName;
  if (!(board in boards)) {
    board = defaultBoardName;
  }
  setBoard(board);
}

// Remove the /s/foobar suffix of the URL, and replace it with a bare /play/ URL
// instead.
function clearShareURL() {
  let parts = shareURLRegexp.exec(document.location.pathname);
  if (parts) {
    let pathname = document.location.pathname;
    pathname = pathname.substring(0, pathname.length-parts[0].length+1);
    history.replaceState({}, '', pathname);
  }
}

function setupShareButton() {
  let modalElement = document.querySelector('#shareModal');
  let input = modalElement.querySelector('input');
  let button = modalElement.querySelector('button.copy-to-clipboard');

  let copiedTooltip = new bootstrap.Tooltip(input, {
    title: 'Copied to clipboard!',
    trigger: 'manual',
    placement: 'bottom',
  });

  let buttonTooltip = bootstrap.Tooltip.getOrCreateInstance(button);

  document.querySelector('#btn-share').addEventListener('click', async e => {
    e.preventDefault();

    // Share modal with 'loading...' input.
    input.value = '';
    button.disabled = true;
    let modal = new bootstrap.Modal(modalElement, {});
    modal.show();

    // Save the current code snippet.
    let id = await simulator.share();

    // Update the modal and URL.
    let url;
    if (shareURLRegexp.test(document.location.pathname)) {
      url = new URL(id, document.location.href);
    } else {
      url = new URL('s/' + id, document.location.href);
    }
    history.replaceState({}, '', url);
    input.value = url;
    button.disabled = false;
  });

  button.addEventListener('click', e => {
    e.preventDefault();
    navigator.clipboard.writeText(input.value);
    input.focus();

    // Hide the 'Copy to clipboard' tooltip on the button.
    buttonTooltip.hide();

    // Show a new tooltip on the input element.
    copiedTooltip.show();
    setTimeout(() => {
      copiedTooltip.hide();
    }, 1500);
  });

  input.addEventListener('focus', e => {
    input.select();
  })
}

// Initialize the playground.
document.addEventListener('DOMContentLoaded', async function(e) {
  editor = new Editor(document.querySelector('.playground-editor'));

  simulator = new Simulator({
    root: document.querySelector('#output'),
    editor: editor,
    firmwareButton: document.querySelector('#btn-flash'),
    baseURL: new URL('/playground/', document.baseURI),
    apiURL: PLAYGROUND_API,
    saveState: () => {
      // Strip /s/* URL suffix.
      clearShareURL();
      localStorage.tinygo_playground_state = JSON.stringify(simulator.schematic.state);
      console.log(simulator.schematic.state);
    },
  });

  setupShareButton();

  // Update the drop down list of boards.
  updateBoards();

  // Now see which editor state we should load (depending on various things).
  loadSavedState();
})
