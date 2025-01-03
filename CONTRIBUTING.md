# Contribution Guidelines

Thank you for considering contributing to the QuoteHub library! Contributions are what make this library better for everyone. If you have a quote you'd like to add, follow these simple steps to include it in the library.

---

## Adding a New Quote

Quotes in QuoteHub are stored in the `quotes.yaml` file. To add a new quote, follow these steps:

### Step 1: Fork the Repository
1. Go to the [QuoteHub repository](https://github.com/kashifkhan0771/quotehub).
2. Click the "Fork" button in the top-right corner to create your own copy of the repository.
3. Clone your forked repository to your local machine:
   ```bash
   git clone https://github.com/kashifkhan0771/quotehub.git
   ```
4. Navigate to the repository directory:
   ```bash
   cd quotehub
   ```

### Step 2: Open `quotes.yaml`
1. Locate the `quotes.yaml` file in the root of the repository.
2. Open it in your favorite text editor.

### Step 3: Add Your Quote
Each quote in `quotes.yaml` follows this format:

```yaml
- text: "Your quote here."
  author: "Author Name"
  category: "Category Name"
  tags:
    - tag1
    - tag2
  year: 2023
```

- **`text`**: The actual quote text.
- **`author`**: The name of the person who said or wrote the quote.
- **`category`**: A category that best describes the quote (e.g., "Inspiration", "Life").
- **`tags`**: A list of keywords associated with the quote for easier filtering.
- **`year`**: The year the quote was created or published (optional).

#### Example:

```yaml
- text: "The only way to do great work is to love what you do."
  author: "Steve Jobs"
  category: "Inspiration"
  tags:
    - work
    - passion
    - success
  year: 2005
```

### Step 4: Run `make quotes`
After updating `quotes.yaml`, regenerate the embedded quotes file:

1. Ensure you have `make` installed on your system.
2. Run the following command in your terminal:
   ```bash
   make quotes
   ```
3. This will generate a new `quotes_data.go` file, embedding the updated `quotes.yaml` content into the library.


### Step 5: Commit and Push Your Changes
1. Commit your updated `quotes.yaml` and the regenerated `quotes_data.go` file:
   ```bash
   git add quotes.yaml quotes_data.go
   git commit -m "Add new quote by [Author Name]"
   ```
2. Push your changes to your forked repository:
   ```bash
   git push origin main
   ```
3. Submit a pull request to the original QuoteHub repository:
   - Go to your forked repository on GitHub.
   - Click "Pull Requests" and then "New Pull Request".
   - Provide a clear description of the changes and submit your pull request.

---

## Need Help?
If you encounter any issues or have questions, feel free to [open an issue](https://github.com/kashifkhan0771/quotehub/issues) on GitHub or contact the maintainers.

Happy quoting!
