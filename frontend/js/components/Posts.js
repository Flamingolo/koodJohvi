export async function fetchPosts(){
    const response = await fetch('http://localhost:8080/posts');
    if (!response.ok){
        throw new Error('Failed to fetch posts');
    }
    return response.json();
}

export async function likePost(postID){
    const response = await fetch(`http://localhost:8080/posts/${postID}/like`, {
        method: 'POST'
    });
    if (!response.ok){
        throw new Error('Failed to like post');
    }
    return response.json();
}

export async function dislikePost(postID) {
    const respone = await fetch(`http://localhost:8080/posts/${postID}/dislike`, {
        method: 'POST'
    });
    if (!response.ok){
        throw new Error('Failed to dislike post');
    }
    return response.json();
}

export function createPostElement(post){
    const postElement = document.createElement('div');
    postElement.innerHTML = `
        <div>
            <h3>${post.title}</h3>
            <p>${post.content}</p>
            <p>Author: ${post.user_id}</p>
            <p>Likes: ${post.likes - post.dislikes}</p>
            <button data-id="${post.id}" class="like-button">Like</button>
            <button data-id="${post.id}" class="dislike-button">Dislike</button>
        </div>
    `;

    postElement.querySelector('.like-button').addEventListener('click', async () => {
       try {
        await likePost(post.id);
        post.likes++
        postElement.querySelector('p:nth-of-type(4)').textContent = `Likes: ${post.likes - post.dislikes}`;
       } catch (error) {
        console.error('Error:', error);
        alert('Failed to like post');
       }
    });

    postElement.querySelectorAll('.dislike-button').addEventListener('click', async () => {
        try {
            await dislikePost(post.id);
            post.likes--;
            postElement.querySelector('p:nth-of-type(4)').textContent = `Likes: ${post.likes - post.dislikes}`;
        } catch (error) {
            console.error('Error:', error);
            alert('Failed to dislike post');
        }
    });

    return postElement;
}

export default async function Posts(){
    const postSection = document.createElement('section');
    postSection.innerHTML = `
        <h2>Posts</h2>
        <div id="posts"></div>
    `;

    try {
        const posts = await fetchPosts();
        const postList = postSection.querySelector('#posts');
        posts.forEach(post => {
            const postElement = createPostElement(post);
            postList.appendChild(postElement);
        });
    } catch (error) {
        alert('Failed to fetch posts');
    }

    return postSection;
}