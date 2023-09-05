# openAPI schema part0 for dialogflowcx.cnrm.cloud.google.com.v1alpha1.DialogflowCXAgent

## schema

```yaml
properties:
  apiVersion:
    description: 'apiVersion defines the versioned schema of this representation of
      an object. Servers should convert recognized schemas to the latest internal
      value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources'
    type: string
  kind:
    description: 'kind is a string value representing the REST resource this object
      represents. Servers may infer this from the endpoint the client submits requests
      to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
    type: string
  metadata:
    type: object
  spec:
    properties:
      avatarUri:
        description: The URI of the agent's avatar. Avatars are used throughout the
          Dialogflow console and in the self-hosted Web Demo integration.
        type: string
      defaultLanguageCode:
        description: |-
          Immutable. The default language of the agent as a language tag. [See Language Support](https://cloud.google.com/dialogflow/cx/docs/reference/language)
          for a list of the currently supported language codes. This field cannot be updated after creation.
        type: string
      description:
        description: The description of this agent. The maximum length is 500 characters.
          If exceeded, the request is rejected.
        type: string
      displayName:
        description: The human-readable name of the agent, unique within the location.
        type: string
      enableSpellCorrection:
        description: Indicates if automatic spell correction is enabled in detect
          intent requests.
        type: boolean
      enableStackdriverLogging:
        description: Determines whether this agent should log conversation queries.
        type: boolean
      location:
        description: |-
          Immutable. The name of the location this agent is located in.

          ~> **Note:** The first time you are deploying an Agent in your project you must configure location settings.
           This is a one time step but at the moment you can only [configure location settings](https://cloud.google.com/dialogflow/cx/docs/concept/region#location-settings) via the Dialogflow CX console.
           Another options is to use global location so you don't need to manually configure location settings.
        type: string
      projectRef:
        description: The project that this resource belongs to.
        oneOf:
        - not:
            required:
            - external
          required:
          - name
        - not:
            anyOf:
            - required:
              - name
            - required:
              - namespace
          required:
          - external
        properties:
          external:
            description: 'Allowed value: The `name` field of a `Project` resource.'
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      resourceID:
        description: Immutable. Optional. The service-generated name of the resource.
          Used for acquisition only. Leave unset to create a new resource.
        type: string
      securitySettings:
        description: 'Name of the SecuritySettings reference for the agent. Format:
          projects/<Project ID>/locations/<Location ID>/securitySettings/<Security
          Settings ID>.'
        type: string
      speechToTextSettings:
        description: Settings related to speech recognition.
        properties:
          enableSpeechAdaptation:
            description: Whether to use speech adaptation for speech recognition.
            type: boolean
        type: object
      supportedLanguageCodes:
        description: The list of all languages supported by this agent (except for
          the default_language_code).
        items:
          type: string
        type: array
      timeZone:
        description: |-
          The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
          Europe/Paris.
        type: string
    required:
    - defaultLanguageCode
    - displayName
    - location
    - projectRef
    - timeZone
    type: object
  status:
    properties:
      conditions:
        description: Conditions represent the latest available observation of the
          resource's current state.
        items:
          properties:
            lastTransitionTime:
              description: Last time the condition transitioned from one status to
                another.
              type: string
            message:
              description: Human-readable message indicating details about last transition.
              type: string
            reason:
              description: Unique, one-word, CamelCase reason for the condition's
                last transition.
              type: string
            status:
              description: Status is the status of the condition. Can be True, False,
                Unknown.
              type: string
            type:
              description: Type is the type of the condition.
              type: string
          type: object
        type: array
      name:
        description: The unique identifier of the agent.
        type: string
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      startFlow:
        description: 'Name of the start flow in this agent. A start flow will be automatically
          created when the agent is created, and can only be deleted by deleting the
          agent. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent
          ID>/flows/<Flow ID>.'
        type: string
    type: object
required:
- spec
type: object

```