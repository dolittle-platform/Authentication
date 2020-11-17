const params = new URLSearchParams(location.search);

$.get('/.ory/kratos/public/sessions/whoami',
    (data) => {
        const tenants = data.identity.traits.tenants;

        const container = $('#container');

        for (const tenant of tenants) {
            const form = $('<form>', {
                action: 'action',
                method: 'POST'
            });
            form.append($('<input>', {
                name: 'tenant',
                type: 'submit',
                value: tenant
            }));
            container.append(form);
        }
    }
);
