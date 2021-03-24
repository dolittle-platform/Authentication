const identityProviders = [
    {
        ID: 'sample',
        Display: 'sample',
        ImageURL: 'https://via.placeholder.com/60'
    }
];

const availableTenants = [
    {
        ID: 'dolittle',
        Display: 'dolittle'
    },
    {
        ID: 'tenant-a',
        Display: 'tenant-a'
    },
    {
        ID: 'tenant-b',
        Display: 'tenant-b'
    }
];

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
                        <a href="/restart"/>Click here to do it all again</a>
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
    res.send({
        ID: req.query.id,
        Forced: false,
        FormCSRFToken: 'csrf-form-token',
        FormSubmitAction: {
            Scheme: 'http',
            Opaque: '',
            User: null,
            Host: 'localhost:8080',
            Path: '/.auth/self-service/methods/oidc/auth/authentication-id',
            RawPath: '',
            ForceQuery: false,
            RawQuery: '',
            Fragment: '',
            RawFragment: ''
        },
        FormSubmitMethod: 'POST',
        Providers: identityProviders,
    });
});

app.post('/.auth/self-service/methods/oidc/auth/authentication-id', (req, res) => {
    console.log('Authenticating with external authority', req.body.provider)
    res.redirect('/.auth/select-tenant?login_challenge=1234');
});

app.get('/.auth/self-service/tenant/flows', (req, res) => {
    res.send({
        ID: req.query.login_challenge,
        FormSubmitAction: {
            Scheme: 'http',
            Opaque: '',
            User: null,
            Host: 'localhost:8080',
            Path: '/.auth/self-service/tenant/select',
            RawPath: '',
            ForceQuery: false,
            RawQuery: '',
            Fragment: '',
            RawFragment: ''
        },
        FormSubmitMethod: 'POST',
        User: {
            Subject: 'subject',
            Tenants: availableTenants,
        }
    });
});

app.post('/.auth/self-service/tenant/select', (req, res) => {
    console.log('Logging in with tenant', req.body.tenant);
    res.cookie('logged_in', 'yes', { httpOnly: true, sameSite: 'strict' });
    res.redirect('/');
});

app.get('/restart', (_, res) => {
    res.cookie('logged_in', 'no', { httpOnly: true, sameSite: 'strict' });
    res.redirect('/');
});

app.use('/', proxy('http://localhost:8091'));

app.listen(8080, () => {
    console.log('Listening at http://localhost:8080');
});
