const params = new URLSearchParams(location.search);

$.get('/.ory/kratos/public/sessions/whoami',
    (data) => {
        const tenants = data.identity.traits.tenants;

        const container = $('#container');

        for (const tenant of tenants) {
            const form = $('<form>', {
                action: 'http://localhost:8080/.auth/selected-tenant/',
                method: 'POST'
            });
            form.append($('<input>', {
                name: 'login_challenge',
                type: 'hidden',
                value: params.get('login_challenge')
            }));
            form.append($('<input>', {
                name: 'tenant',
                type: 'submit',
                value: tenant
            }));
            container.append(form);
        }
    }
);
