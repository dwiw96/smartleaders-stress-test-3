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
});
