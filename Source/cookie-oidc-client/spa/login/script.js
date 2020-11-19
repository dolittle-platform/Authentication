const params = new URLSearchParams(location.search);

$.get(
    '/.ory/kratos/public/self-service/login/flows?id='+params.get('flow'),
    (data) => {
        console.log('we gucci from kratos', data);
        const container = $('#container');
        for (const methodId in data.methods) {
            const config = data.methods[methodId].config;
            
            const form = $('<form>', {
                action: config.action,
                method: config.method
            });
            for (const field of config.fields) {
                form.append($('<input>', {
                    name: field.name,
                    type: field.type,
                    value: field.value,
                    required: field.required
                }));
            }

            container.append(form);
            console.log(config);
        }

    });

