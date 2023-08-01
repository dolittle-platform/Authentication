const identityProviders = [
    {
        id: 'sample',
        display: 'sample',
        imageURL: 'https://via.placeholder.com/60',
    },
    {
        id: 'sample-2',
        display: 'sample 2',
        imageURL: 'https://via.placeholder.com/60',
    },
    {
        id: 'sample-3',
        display: 'sample no tenant',
        imageURL: 'https://via.placeholder.com/60'
    },
];

const availableTenants = [
    {
        id: 'dolittle',
        display: 'dolittle',
    },
    {
        id: 'tenant-a',
        display: 'tenant-a',
    },
    {
        id: 'tenant-b',
        display: 'tenant-b'
    },
];

const responseDelay = 1000;

const app = require('express')();
const cookieParser = require('cookie-parser');
const bodyParser = require('body-parser').urlencoded;
const proxy = require('express-http-proxy');

app.use(cookieParser());
app.use(bodyParser({ extended: false }));

app.get('/', (req, res) => {
    if (req.cookies.logged_in == 'yes') {
        res.send(`
            <!DOCTYPE html>
            <html>
                <head>
                    <title>Logged in</title>
                </head>
                <body>
                    <h1>You are logged in!</h1>
                    <div>
                        <a href="/.auth/cookies/logout"/>Click here to logout and do it all again</a>
                    </div>
                    <div>
                        <a href="/.auth/error?correlation=error-id"/>Click here to see what an error would look like</a>
                    </div>
                </body>
            </html>
        `);
    } else {
        console.log('Starting login flow...');
        res.redirect('/.auth/select-provider?flow=1234')
    }
});

app.get('/.auth/self-service/login/flows', (req, res) => {
    setTimeout(() => {
        res.send({
            id: req.query.id,
            refresh: false,
            form: {
                csrfToken: 'csrf-form-token',
                submitAction: 'http://localhost:8080/.auth/self-service/methods/oidc/auth/authentication-id',
                submitMethod: 'POST',
            },
            providers: identityProviders,
        });
    }, responseDelay);
});

app.post('/.auth/self-service/methods/oidc/auth/authentication-id', (req, res) => {
    console.log('Authenticating with external authority', req.body.provider)
    if (req.body.provider == 'sample-3') {
        res.redirect('/.auth/no-tenant');
    } else {
        res.redirect('/.auth/select-tenant?login_challenge=1234');
    }
});

app.get('/.auth/self-service/tenant/flows', (req, res) => {
    setTimeout(() => {
        res.send({
            id: req.query.login_challenge,
            form: {
                submitAction: 'http://localhost:8080/.auth/self-service/tenant/select',
                submitMethod: 'POST',
            },
            user: {
                subject: 'subject',
                tenants: availableTenants,
            }
        });
    }, responseDelay);
});

app.post('/.auth/self-service/tenant/select', (req, res) => {
    console.log('Logging in with tenant', req.body.tenant);
    res.cookie('logged_in', 'yes', { httpOnly: true, sameSite: 'strict' });
    res.redirect('/');
});

app.get('/.auth/cookies/logout', (_, res) => {
    res.cookie('logged_in', 'no', { httpOnly: true, sameSite: 'strict' });
    res.redirect('/.auth/logged-out');
});

app.use('/', proxy('http://localhost:8091'));

app.listen(8080, () => {
    console.log('Listening at http://localhost:8080');
});
