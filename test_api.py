import requests
import uuid

base_url = "http://localhost:3000/api"

def test_endpoints():
    print("Testing Endpoints...")
    
    # Register a new user
    user_id = str(uuid.uuid4())
    print(f"\n--- Testing User Routes ---")
    resp = requests.post(f"{base_url}/users", json={
        "id": user_id,
        "name": "Test User",
        "email": "test@example.com",
        "password": "password"
    })
    print("POST /users:", resp.status_code, resp.text)
    
    # Login
    resp = requests.post(f"{base_url}/users/_login", json={
        "email": "test@example.com",
        "password": "password"
    })
    print("POST /users/_login:", resp.status_code, resp.text)
    token = ""
    if resp.status_code == 200:
        token = resp.json().get("data", {}).get("token", "")
    
    headers = {"Authorization": token}
    
    # Categories
    print(f"\n--- Testing Category Routes ---")
    resp = requests.get(f"{base_url}/categories", headers=headers)
    print("GET /categories:", resp.status_code, resp.text[:100])
    
    # Courses
    print(f"\n--- Testing Course Routes ---")
    resp = requests.get(f"{base_url}/courses", headers=headers)
    print("GET /courses:", resp.status_code, resp.text[:100])

    # Roles
    print(f"\n--- Testing Role Routes ---")
    resp = requests.get(f"{base_url}/roles", headers=headers)
    print("GET /roles:", resp.status_code, resp.text[:100])

if __name__ == "__main__":
    test_endpoints()
