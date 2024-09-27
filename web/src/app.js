document.addEventListener('DOMContentLoaded', () => {
    const form = document.getElementById('form');
    
    form.addEventListener('submit', async (e) => {
        e.preventDefault(); // Prevent the form from submitting the traditional way

        // Collect the form data
        const formData = new FormData(form);

        // Convert form data to a JSON object
        const data = {
            fullname: formData.get('fullname'),
            physical_address: formData.get('address'),
            movie_rented: formData.get('movie'),
            salutation: formData.get('salutation')
        };

        console.log(data)

        try {
            // Post the data to the API endpoint
            const response = await fetch('http://localhost:8080/api/rent_book', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data)
            });

            if (response.ok) {
                const result = await response.json();
                alert('Form submitted successfully');
                console.log(result);
            } else {
                alert('Failed to submit form');
            }
        } catch (error) {
            console.error('Error:', error);
            alert('Error submitting form');
        }
    });

    document.getElementById('goToTable').addEventListener('click', async function() {
        try {
            // Fetch data from the API
            const response = await fetch('http://localhost:8080/api/list_of_rented_books');
            if (!response.ok) throw new Error('Network response was not ok');
            
            const data = await response.json();

            // Check if the result is success and value is present
            if (data.result === 'success' && data.value) {
                // Save the data in localStorage for use in table.html
                localStorage.setItem('rentedBooks', JSON.stringify(data.value));

                // Redirect to table.html
                window.location.href = 'table.html';
            } else {
                alert('Failed to fetch rented books data');
            }
        } catch (error) {
            console.error('Error fetching data:', error);
            alert('Error fetching rented books');
        }
    });
});
