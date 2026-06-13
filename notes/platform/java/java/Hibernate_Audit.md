# Hibernate Auditing

@EnableJpaAuditing
@EntityListeners { AuditListener.class }

    @OneToOne( ?
        mappedBy = "post",
        cascade = CascadeType.ALL,
        orphanRemoval = true,
        fetch = FetchType.LAZY
    )


    @OneToMany( ?
        mappedBy = "post",
        cascade = CascadeType.ALL,
        orphanRemoval = true
    )

    @GeneratorType( ?
        type = LoggedUserGenerator.class,
        when = GenerationTime.INSERT 
        when = GenerationTime.ALWAYS
    )

    @Embedded ?
    private Audit audit;

### Entity Event

@PrePersist: Executed before the entity manager persist operation is actually executed or cascaded. This call is synchronous with the persist operation.
@PreRemove:  Executed before the entity manager remove operation is actually executed or cascaded. This call is synchronous with the remove operation.
@PostPersist:    Executed after the entity manager persist operation is actually executed or cascaded. This call is invoked after the database INSERT is executed.
@PostRemove: Executed after the entity manager remove operation is actually executed or cascaded. This call is synchronous with the remove operation.
@PreUpdate:  Executed before the database UPDATE operation.
@PostUpdate: Executed after the database UPDATE operation.
@PostLoad:   Executed after an entity has been loaded into the current persistence context or an entity has been refreshed.

### EventListener Order

Once event is emmited, listener will be executed in this order.

1. Entity listeners for the superclasses (highest first)
2. Entity Listeners for the entity
3. Callbacks of the superclasses (highest first)
4. Callbacks of the entity
