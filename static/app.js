document.addEventListener("DOMContentLoaded", () => {
  // 1. Select all code blocks (pre tags)
  const codeBlocks = document.querySelectorAll(".post pre");

  codeBlocks.forEach((block) => {
    // 2. Create the button
    const button = document.createElement("button");
    button.className = "copy-btn";
    button.innerText = "Copy";

    // 3. Add the Click Event
    button.addEventListener("click", () => {
      // Get the text inside the <code> tag
      const code = block.querySelector("code").innerText;

      // Use the Clipboard API
      navigator.clipboard.writeText(code).then(() => {
        // Visual Feedback
        const originalText = button.innerText;
        button.innerText = "Copied!";

        // Reset after 2 seconds
        setTimeout(() => {
          button.innerText = originalText;
        }, 2000);
      });
    });

    // 4. Append button to the pre block
    block.appendChild(button);
  });
});
