export async function fetchComments(postID) {
    const response = await fetch(`http://localhost:8080/posts/${postID}/comments`)
    if (!response.ok){
        throw new Error('Failed to fetch comments')
    }
    return response.json()
}

export async function likeComment(commentID) {
    const response = await fetch(`http://localhost:8080/comments/${commentID}/like`, {
        method: 'POST',
    })
    if (!response.ok){
        throw new Error('Failed to like comment')
    }

    return response.json()
}

export async function dislikeComment(commentID) {
    const response = await fetch(`http://localhost:8080/comments/${commentID}/dislike`, {
        method: 'POST',
    });
    if (!response.ok){
        throw new Error('Failed to dislike comment')
    }
    return response.json();
}

export function createCommentElement(comment){
    const commentElement = document.createElement('div');
    commentElement.innerHTML = `
        <p>${comment.conent}</p>
        <p>Author: ${comment.user_id}</p>
        <p>Likes: ${comment.likes - comment.dislikes}</p>
        <button data-id=${comment.id} class="like-button">Like</button>
        <button data-id=${comment.id} class="dislike-button">Dislike</button>
    `;

    commentElement.querySelector(".like-button").addEventListener('click', async () => {
        try {
            await likeComment(comment.id);
            comment.likes++;
            commentElement.querySelector('p:nth-of-type(3)').innerText = `Likes: ${comment.likes - comment.dislikes}`;
        } catch (error) {
            console.error('Error:', error);
            alert('Failed to like comment');
        }
    });

    return commentElement;
}

export default async function Comments(postID){
    const commentSection = document.createElement('section');
    commentSection.innerHTML = `
        <h2>Comments</h2>
        <div id="comments-list"></div>
    `;

    try {
        const comments = await fetchComments(postID);
        const commentsList = commentSection.querySelector('#comments-list');
        comments.forEach(ocmment => {
            const commentElement = createCommentElement(comment);
            commentsList.appendChild(commentElement);
        });
    } catch (erorr){
        alert('Failed to load comments');
    }

    return commentSection
}
