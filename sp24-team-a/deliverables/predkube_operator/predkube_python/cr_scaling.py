from kubernetes import client, config

GROUP = 'webapp.test.com'
VERSION = 'v1'
NAMESPACE = 'default'
PLURAL = 'scalings'

def get_existing_scaling(api_instance, namespace, deployment_name):
    try:
        resources = api_instance.list_namespaced_custom_object(
            group=GROUP,
            version=VERSION,
            namespace=NAMESPACE,
            plural=PLURAL
        )
        for item in resources.get('items', []):
            if item['spec']['targetDeployment'] == deployment_name and item['spec']['targetNamespace'] == namespace:
                return item
    except Exception as e:
        print("No Existing Scaling")
        return None
    return None

def update_scaling_resource(api_instance, name, new_replica_count):
    try:
        resource = api_instance.get_namespaced_custom_object(
            group=GROUP,
            version=VERSION,
            namespace=NAMESPACE,
            plural=PLURAL,
            name=name
        )
        if resource['spec']['targetReplicaCount'] != new_replica_count:
            resource['spec']['targetReplicaCount'] = new_replica_count
            api_response = api_instance.patch_namespaced_custom_object(
                group=GROUP,
                version=VERSION,
                namespace=NAMESPACE,
                plural=PLURAL,
                name=name,
                body=resource
            )
            print("Scaling Updated.")
        else:
            print("No Update.")
    except Exception as e:
        print("Failed to update.")

def create_scaling_resource(deployment, namespace, replica_count, context='kind-test'):
    config.load_kube_config(context=context)

    api_instance = client.CustomObjectsApi()

    existing = get_existing_scaling(api_instance, namespace, deployment)
    if existing is not None:
        update_scaling_resource(api_instance, existing['metadata']['name'], replica_count)
    else:
        scaling_name = namespace + "-" + deployment + "-scaling"

        cr_body = {
            "apiVersion": "webapp.test.com/v1",
            "kind": "Scaling",
            "metadata": {
                "name": scaling_name
            },
            "spec": {
                "targetDeployment": deployment,
                "targetNamespace": namespace,
                "targetReplicaCount": replica_count
            }
        }

        try:
            api_response = api_instance.create_namespaced_custom_object(
                group=GROUP,
                version=VERSION,
                namespace=NAMESPACE,
                plural=PLURAL,
                body=cr_body
            )
            print("Scaling Created.")
        except client.ApiException as e:
            print("Failed to create, ", e)

# create_scaling_resource('test-app', 'default', 3)
