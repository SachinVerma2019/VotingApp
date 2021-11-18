import { useState } from 'react';
import { useHistory } from 'react-router-dom';
import axios from 'axios';
import { useToken } from '../auth/useToken';

export const LogInPage = () => {
    const [token, setToken] = useToken();

    const [errorMessage, setErrorMessage] = useState('');

    const [emailValue, setEmailValue] = useState('');
    const [passwordValue, setPasswordValue] = useState('');

    const history = useHistory();

    const onLogInClicked = async () => {
        var bodyFormData = new FormData();
        bodyFormData.append('email', emailValue);
        bodyFormData.append('password', passwordValue);
        const response = await axios({
            method: 'post',
            url: '/user/login',
            data: bodyFormData,
            headers: {'Content-Type': 'multipart/form-data' }
            // headers: {
            //     'Content-Type': `multipart/form-data; boundary=${form._boundary}`,
            // },
        });

        const { token } = response.data;
        localStorage.setItem('userId', response.data.id);
        localStorage.setItem('userName', response.data.name);
        setToken(token);
        history.push('/');
    }

    return (
        <div className="content-container">
            <h1>Log In</h1>
            {errorMessage && <div className="fail">{errorMessage}</div>}
            <input
                value={emailValue}
                onChange={e => setEmailValue(e.target.value)}
                placeholder="someone@gmail.com" />
            <input
                type="password"
                value={passwordValue}
                onChange={e => setPasswordValue(e.target.value)}
                placeholder="password" />
            <hr />
            <button
                disabled={!emailValue || !passwordValue}
                onClick={onLogInClicked}>Log In</button>
            <button onClick={() => history.push('/forgot-password')}>Forgot your password?</button>
            <button onClick={() => history.push('/signup')}>Don't have an account? Sign Up</button>
        </div>
    );
}